package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/beeploop/our-srts/internal/config"
)

func main() {
	cfg := config.Load()

	filename := "test_file.txt"
	path := filepath.Join(cfg.UPLOAD_DIR, filename)

	data := []byte("hello world\n")

	if err := os.WriteFile(path, data, 0777); err != nil {
		log.Fatalf("could not write test file: %s\n", err.Error())
	}

	log.Println("written test file")
}
