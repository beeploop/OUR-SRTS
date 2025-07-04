package main

import (
	"context"
	"flag"
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

	mode := flag.String("mode", "", "seeding mode (admin, program, student, documentType)")
	source := flag.String("source", "", "source file")
	limit := flag.Int("limit", 0, "limit number of students inserted")

	flag.Parse()

	ctx := context.Background()

	if *source == "" {
		log.Fatalf("invalid source file")
	}

	switch *mode {
	case "admin":
		adminRepo := repositories.NewAdminRepository(db)
		adminSeeder := seeder.NewAdminSeeder(*source, adminRepo)

		if err := adminSeeder.Execute(ctx); err != nil {
			log.Fatalf("failed to seed admins: %s\n", err.Error())
		}
	case "program":
		programRepo := repositories.NewProgramRepository(db)
		programSeeder := seeder.NewProgramSeeder(*source, programRepo)

		if err := programSeeder.Execute(ctx); err != nil {
			log.Fatalf("failed to seed programs: %s\n", err.Error())
		}

	case "student":
		studentRepo := repositories.NewStudentRepository(db)
		studentSeeder := seeder.NewStudentSeeder(*source, studentRepo)

		if err := studentSeeder.Execute(ctx, limit); err != nil {
			log.Fatalf("failed to seed students: %s\n", err.Error())
		}

	case "documentType":
		documentTypeRepo := repositories.NewDocumentTypeRepository(db)
		documentTypeSeeder := seeder.NewDocumentTypeSeeder(*source, documentTypeRepo)

		if err := documentTypeSeeder.Execute(ctx); err != nil {
			log.Fatalf("failed to seed document types: %s\n", err.Error())
		}

	default:
		log.Fatalf("Invalid mode selected")
	}

	log.Printf("seeding complete\n")
}
