package main

import (
	"net/http"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/server"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.Info("Loading configuration...")
	err := config.Initialize()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Initializing database...")
	err = store.Initialize()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Initializing database tables...")
	store.InitializeTables()
}

func main() {
	logrus.Info("Starting server...")
	server.NewServer()

	logrus.Infof("Server listening on %s%s\n", config.Env.LocalAddr, config.Env.Port)
	err := http.ListenAndServe(config.Env.Port, server.Router)
	if err != nil {
		logrus.Fatal(err)
	}
}
