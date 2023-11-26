package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/BeepLoop/registrar-digitized/config"
)

func InitSession() cookie.Store {
	store := cookie.NewStore([]byte("my-session-secret"))
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Domain:   config.Env.LocalAddr,
		MaxAge:   60 * 60 * 24,
	})

	return store
}
