package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/utils"
	"strconv"
	"strings"
)

var (
	ErrInvalidTokenFormat = errors.New("невалидный формат токена")
	ErrInvalidTokenTime   = errors.New("невалидное время токена")
	ErrTokenExpired       = errors.New("время жизни токена закончилось")
	ErrInvalidMAC         = errors.New("внутренняя ошибка сервиса")
)

func (csrf *CscfUsecase) IsValidCSRFToken(session model.Session, token string) (bool, error) {
	if session.SessionID == "" || session.UserID == 0 || token == "" {
		return false, ErrInvalidTokenFormat
	}

	mac, exp, err := parseToken(token)
	if err != nil {
		return false, errors.New("невалидный токен")
	}

	if utils.IsExpired(exp) {
		return false, ErrTokenExpired
	}

	expectedMAC := csrf.generateMAC(csrf.secret, session.SessionID, session.UserID, exp)
	return hmac.Equal(mac, expectedMAC), nil
}

func parseToken(token string) ([]byte, int64, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return nil, 0, ErrInvalidTokenFormat
	}

	exp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, 0, ErrInvalidTokenTime
	}

	mac, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, 0, ErrInvalidMAC
	}

	return mac, exp, nil
}

func (csrf *CscfUsecase) generateMAC(secret, sessionID string, userID uint32, exp int64) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	data := fmt.Sprintf("%s:%d:%d", sessionID, userID, exp)
	mac.Write([]byte(data))

	return mac.Sum(nil)
}
