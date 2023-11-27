package server

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"time"

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

	datetime := time.Now().Format("2006-01-02_15:04:05")
	logFile := "logs/server-log_" + datetime + ".log"

	myFile, err := os.Create(logFile)
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
	Router.Static("/documents", "documents/")
	Router.Static("/fonts", "webfonts/")

	RegisterRoutes()
}
