package main

import (
	"context"
	"log"
	"slices"

	"github.com/beeploop/our-srts/internal/application/seeder"
	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/repositories"
	"github.com/beeploop/our-srts/internal/pkg/utils"
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

	adminsSourceFile := "./seed_files/admins.json"
	adminRepo := repositories.NewAdminRepository(db)
	adminSeeder := seeder.NewAdminSeeder(adminsSourceFile, adminRepo)

	if err := adminSeeder.Execute(ctx); err != nil {
		log.Fatalf("failed to seed admins: %s\n", err.Error())
	}

	programSourceFile := "./seed_files/data.json"
	programRepo := repositories.NewProgramRepository(db)
	programSeeder := seeder.NewProgramSeeder(programSourceFile, programRepo)

	if err := programSeeder.Execute(ctx); err != nil {
		log.Fatalf("failed to seed programs: %s\n", err.Error())
	}

	studentSourceFile := "./seed_files/tagum-trimmed.csv"
	studentRepo := repositories.NewStudentRepository(db)
	studentSeeder := seeder.NewStudentSeeder(studentSourceFile, studentRepo)
	studentLimit := 10

	if err := studentSeeder.Execute(ctx, &studentLimit); err != nil {
		log.Fatalf("failed to seed students: %s\n", err.Error())
	}

	documentTypeNames := []string{
		"photo",
		"birth_certificate",
		"notice_of_admission",
		"usepat",
		"good_moral",
		"form_137",
		"form_138",
		"personal_data_sheet",
		"data_privacy_provision",
		"honorable_dismissal",
		"marriage_certificate",
		"promisory_note",
		"health_state_declaration",
		"medical_certificate",
		"approval_sheet",
		"application_for_graduation",
		"shifters_form",
		"certificate_of_low_income",
		"transcript_of_records",
		"indigency",
		"clearance",
		"leave_of_absence",
		"advance_credit_form",
		"inc_form",
		"subject_validation_form",
		"substitution",
		"affidavit_of_undertaking",
		"nmat",
	}
	documentTypes := slices.AppendSeq(
		make([]entities.DocumentType, 0),
		utils.Map(documentTypeNames, func(documentTypeName string) entities.DocumentType {
			return *entities.NewDocumentType(documentTypeName)
		}),
	)

	documentTypeRepo := repositories.NewDocumentTypeRepository(db)
	documentTypeSeeder := seeder.NewDocumentTypeSeeder(documentTypes, documentTypeRepo)

	if err := documentTypeSeeder.Execute(ctx); err != nil {
		log.Fatalf("failed to seed document types: %s\n", err.Error())
	}

	log.Printf("seeding complete\n")
}
