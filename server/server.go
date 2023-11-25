package server

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/registrar/middleware"
	"github.com/registrar/types"
)

var Router *gin.Engine

func NewServer() {
	gob.Register(types.User{})
	gob.Register([]types.Student{})

	Router = gin.Default()

	sessionStore := InitSession()

	Router.Use(middleware.MimeType)
	Router.Use(sessions.Sessions("user", sessionStore))

	Router.Static("/styles", "views/styles/")
	Router.Static("/scripts", "views/scripts/")

	Router.Static("/public", "assets/public/")

	Router.Static("/documents", "documents/")

	RegisterRoutes()
}
