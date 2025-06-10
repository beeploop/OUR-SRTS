package main

import (
	"context"
	"log"

	"github.com/beeploop/our-srts/internal/application/seeder"
	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Load()

	db, err := persistence.NewMysql(mysql.Config{
		User:                 cfg.DB_USER,
		Passwd:               cfg.DB_PASS,
		Net:                  cfg.DB_NET,
		Addr:                 cfg.DB_HOST + ":" + cfg.DB_PORT,
		DBName:               cfg.DB_NAME,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("could not start db: %s\n", err.Error())
	}
	defer db.Close()

	ctx := context.Background()

	programSourceFile := "./seed_files/data.json"
	programRepo := repositories.NewProgramRepository(db)
	programSeeder := seeder.NewProgramSeeder(programSourceFile, programRepo)

	if err := programSeeder.Execute(ctx); err != nil {
		log.Fatalf("failed to seed programs: %s\n", err.Error())
	}

	log.Printf("seeding complete\n")
}
