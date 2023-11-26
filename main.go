package main

import (
	"log"
	"net/http"

	"github.com/registrar/config"
	"github.com/registrar/server"
	"github.com/registrar/store"
)

func main() {
	err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	err = store.Init()
	if err != nil {
		log.Fatal(err)
	}

	store.InitializeTables()

	server.NewServer()

	log.Printf("Server listening on addr: %s: %s\n", config.Env.LocalAddr, config.Env.Port)
	err = http.ListenAndServe(config.Env.Port, server.Router)
	if err != nil {
		log.Fatal(err)
	}
}
