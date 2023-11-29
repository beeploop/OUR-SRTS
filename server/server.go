package server

import (
	"encoding/gob"
	"io"
	"log"
	"os"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/middleware"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func NewServer() {
	gob.Register(types.User{})
	gob.Register([]types.Student{})

	myFile, err := os.Create("./server-logs.log")
	if err != nil {
		log.Fatal("Error creating log file: ", err)
	}
	gin.DefaultWriter = io.MultiWriter(myFile, os.Stdout)

	if config.Env.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()

	sessionStore := InitSession()

	Router.Use(middleware.MimeType)
	Router.Use(sessions.Sessions("user", sessionStore))

	Router.Static("/styles", "views/styles/")
	Router.Static("/scripts", "views/scripts/")
	Router.Static("/public", "assets/public/")
	Router.Static("/fonts", "webfonts/")

	Router.Static("/download", "release/")

	RegisterRoutes()
}
