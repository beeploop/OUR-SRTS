package session

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

type SessionManager struct {
	store        *sessions.CookieStore
	sessionName  string
	sessionValue string
}

func NewSessionManager(secret []byte) *SessionManager {
	store := sessions.NewCookieStore(secret)

	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int((time.Hour * 24).Seconds()),
	}

	return &SessionManager{
		store:        store,
		sessionName:  "app_session",
		sessionValue: "admin",
	}
}

func (s *SessionManager) SetSession(w http.ResponseWriter, r *http.Request, admin SessionModel) error {
	session, err := s.store.Get(r, s.sessionName)
	if err != nil {
		return err
	}

	session.Values[s.sessionValue] = admin
	return session.Save(r, w)
}

func (s *SessionManager) ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.store.Get(r, s.sessionName)
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

func (s *SessionManager) GetAdmin(r *http.Request) (*SessionModel, bool) {
	session, err := s.store.Get(r, s.sessionName)
	if err != nil {
		return nil, false
	}

	admin, ok := session.Values[s.sessionValue].(SessionModel)
	return &admin, ok
}
