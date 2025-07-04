package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beeploop/our-srts/internal/config"
)

func main() {
	cfg := config.Load()

	files, err := os.ReadDir(cfg.UPLOAD_DIR)
	if err != nil {
		log.Fatalf("could not read UPLOAD_DIR: %s\n", err.Error())
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
