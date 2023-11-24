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

	err = http.ListenAndServe(":3000", server.Router)
	if err != nil {
		log.Fatal(err)
	}
}
