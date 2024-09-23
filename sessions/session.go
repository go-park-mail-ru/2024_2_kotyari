package sessions

import (
	"net/http"

	"2024_2_kotyari/config"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(config.SessionKey))

// GetSession получает сессию
func GetSession(r *http.Request, sessionName string) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}

// SaveSession сохраняет сессию
func SaveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return session.Save(r, w)
}
