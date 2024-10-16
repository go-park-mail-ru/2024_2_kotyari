package user

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	second     = 1
	minute     = 60 * second
	hour       = 60 * minute
	nullTime   = -1
	sessionKey = "user_id"
)

type SessionInterface interface {
	Get(r *http.Request) (*sessions.Session, error)
	Save(w http.ResponseWriter, r *http.Request, session *sessions.Session) error
}

type SessionManager struct {
	sessions sessions.Store
	cfg      sessionConfig
}

func NewSessions() *SessionManager {
	cfg := initSessions()

	return &SessionManager{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
		cfg:      cfg,
	}
}

func newTestSessions() *SessionManager {
	cfg := initTestSession()

	return &SessionManager{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
		cfg:      cfg,
	}
}

func (s *SessionManager) Get(r *http.Request) (*sessions.Session, error) {
	session, err := s.sessions.Get(r, s.cfg.SessionName)

	return session, err
}

func (s *SessionManager) Save(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	return s.sessions.Save(r, w, session)
}

func setSessionOptions(session *sessions.Session, maxAge int) {
	session.Options.MaxAge = maxAge
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteLaxMode
	session.Options.Secure = false
}
