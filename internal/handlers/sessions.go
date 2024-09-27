package handlers

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/config"
	"github.com/gorilla/sessions"
	"net/http"
)

type sessionManager struct {
	sessions sessions.Store
	cfg      config.Session
}

func newSessions() *sessionManager {
	cfg := config.InitSessions()

	return &sessionManager{
		sessions: sessions.NewCookieStore([]byte(cfg.SessionKey)),
		cfg:      cfg,
	}
}

func newTestSessions() *sessionManager {
	cfg := config.InitTestSession()

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
