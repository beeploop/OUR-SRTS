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
	"github.com/sirupsen/logrus"
)

var Router *gin.Engine

func NewServer() {
	gob.Register(types.User{})
	gob.Register([]types.Student{})

	logLevel, err := logrus.ParseLevel(config.Env.LogLevel)
	if err != nil {
		log.Fatal("Error parsing logrus level in the env: ", err)
	}
	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logFile, err := os.OpenFile("./server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating log file: ", err)
	}
	multiwriter := io.MultiWriter(logFile, os.Stdout)
	logrus.SetOutput(multiwriter)

	if config.Env.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.New()

	sessionStore := InitSession()

	Router.Use(gin.Recovery())
	Router.Use(middleware.LogrusMiddleware)
	Router.Use(middleware.MimeType)
	Router.Use(sessions.Sessions("user", sessionStore))

	Router.Static("/styles", "views/styles/")
	Router.Static("/scripts", "views/scripts/")
	Router.Static("/public", "assets/public/")
	Router.Static("/fonts", "webfonts/")

	RegisterRoutes()
}
