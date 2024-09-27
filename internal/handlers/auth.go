package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/db"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"golang.org/x/crypto/argon2"
)

type AuthApp struct {
	db       *db.Users
	sessions *sessionManager
}

func NewApp() *AuthApp {
	return &AuthApp{
		sessions: newSessions(),
	}
}

func validateLogin(w *http.ResponseWriter, creds credsApiRequest) bool {
	return validateEmailAndPassword(w, creds)
}

func validateRegistration(w *http.ResponseWriter, creds credsApiRequest) bool {
	if !validateEmailAndPassword(w, creds) {
		return false
	}

	if creds.Password != creds.RepeatPassword {
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.PasswordsDoNotMatch.Error(),
		})
		return false
	}

	if !isValidUsername(creds.Username) {
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidUsernameFormat.Error(),
		})
		return false
	}

	return true
}

func validateEmailAndPassword(w *http.ResponseWriter, creds credsApiRequest) bool {
	switch {
	case !isValidEmail(creds.Email):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidEmailFormat.Error(),
		})
		return false
	case !isValidPassword(creds.Password):
		writeJSON(*w, http.StatusBadRequest, errs.HTTPErrorResponse{
			ErrorMessage: errs.InvalidPasswordFormat.Error(),
		})
		return false
	}

	return true
}

type Server struct {
	sessions sessions.Store
}

func newAppForTests() *AuthApp {
	return &AuthApp{
		sessions: newTestSessions(),
		db:       db.InitUsersWithData(),
	}
}

type UsernameResponse struct {
	Username string `json:"username"`
}

// Login обрабатывает вход пользователя и создает сессию
// @Summary Логин пользователя
// @Description Проверяет учетные данные пользователя и создает сессию при успешной аутентификации
// @Tags auth
// @Accept json
// @Produce json
// @Param user body db.User true "Данные пользователя"
// @Success 200 {string} string "Успешный вход"
// @Failure 400 {string} string "Неправильный запрос"
// @Failure 401 {string} string "Неправильные учетные данные"
// @Failure 500 {string} string "Ошибка при создании сессии"
// @Router /login [post]
func (a *AuthApp) Login(w http.ResponseWriter, r *http.Request) {
	var creds credsApiRequest

	session, err := a.sessions.Get(r)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.SessionCreationError.Error(),
		})
		return
	}

	if email, isAuthenticated := session.Values["user_id"].(string); isAuthenticated {
		user, _ := a.db.GetUserByEmail(email)
		writeJSON(w, http.StatusOK, UsernameResponse{Username: user.Username})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errs.HTTPErrorResponse{
			ErrorMessage: errs.InternalServerError.Error(),
		})
		return
	}

	if !validateLogin(&w, creds) {
		return
	}

	user, exists := a.db.GetUserByEmail(creds.Email)
	if !exists || !verifyPassword(user.Password, creds.Password) {
		writeJSON(w, http.StatusUnauthorized, errs.HTTPErrorResponse{
			ErrorMessage: errs.UnauthorizedCredentials.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, UsernameResponse{Username: user.Username})
}

const (
	timeCost    uint32 = 1         // Время обработки (количество итераций)
	memoryCost  uint32 = 64 * 1024 // Память, используемая Argon2 (в KB)
	parallelism uint8  = 4         // Количество параллельных потоков
	keyLength   uint32 = 32        // Длина генерируемого ключа
	saltLength  int    = 16        // Длина соли в байтах
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)
	fullHash := append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(fullHash)
}

// Разделение соли и хеша
func splitSaltAndHash(saltHashBase64 string) ([]byte, []byte, error) {
	saltHash, err := base64.RawStdEncoding.DecodeString(saltHashBase64)
	if err != nil {
		return nil, nil, err
	}

	salt := saltHash[:saltLength] // Первые saltLength байт — это соль
	hash := saltHash[saltLength:] // Остальное — это хеш
	return salt, hash, nil
}

func verifyPassword(storedSaltHashBase64, password string) bool {
	salt, storedHash, err := splitSaltAndHash(storedSaltHashBase64)
	if err != nil {
		return false
	}
	computedHash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLength)

	return string(storedHash) == string(computedHash)
}

// Logout очищает куки и завершает сессию
// @Summary Логаут пользователя
// @Description Завершает сессию пользователя, очищая куки и удаляя все значения из сессии
// @Tags auth
// @Produce json
// @Success 200 {string} string "Вы успешно вышли"
// @Failure 401 {string} string "Пользователь не авторизован"
// @Failure 500 {string} string "Ошибка при завершении сессии"
// @Router /logout [post]
func (a *AuthApp) Logout(w http.ResponseWriter, r *http.Request) {
	// Получаем сессию из запроса
	session, err := a.sessions.Get(r)
	if err != nil {
		http.Error(w, errs.UnauthorizedMessage.Error(), http.StatusUnauthorized)
		return
	}
	// Очищаем значения сессии, создавая новую пустую мапу
	session.Values = make(map[interface{}]interface{})

	// Устанавливаем время жизни сессии в -1, что означает, что сессия будет завершена
	session.Options.MaxAge = -1

	// Сохраняем изменения сессии
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, errs.LogoutError.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
