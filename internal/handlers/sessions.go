package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type sessionManager struct {
	sessions sessions.Store
	cfg      session
}

func newSessions() *sessionManager {
	cfg := initSessions()

	return &sessionManager{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
		cfg:      cfg,
	}
}

func newTestSessions() *sessionManager {
	cfg := initTestSession()

	return &sessionManager{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
		cfg:      cfg,
	}
}

func (s *sessionManager) Get(r *http.Request) (*sessions.Session, error) {
	session, err := s.sessions.Get(r, s.cfg.SessionName)

	return session, err
}

func (s *sessionManager) Save(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	return s.sessions.Save(r, w, session)
}
