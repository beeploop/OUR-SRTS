package main

import (
	"log"
	"net/http"

	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/server"
	"github.com/BeepLoop/registrar-digitized/store"
)

func init() {
	err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	err = store.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	store.InitializeTables()
}

func main() {
	server.NewServer()

	log.Printf("Server listening on %s%s\n", config.Env.LocalAddr, config.Env.Port)
	err := http.ListenAndServe(config.Env.Port, server.Router)
	if err != nil {
		log.Fatal(err)
	}
}
