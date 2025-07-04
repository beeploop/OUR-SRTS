package seeder

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type StudentSeeder struct {
	sourceFile  string
	studentRepo repositories.StudentRepository
}

func NewStudentSeeder(sourceFile string, repo repositories.StudentRepository) *StudentSeeder {
	return &StudentSeeder{
		sourceFile:  sourceFile,
		studentRepo: repo,
	}
}

func (s *StudentSeeder) Execute(ctx context.Context, limit *int) error {
	records, err := s.readFile()
	if err != nil {
		return err
	}

	for i, record := range records {
		if limit != nil && *limit != 0 && i+1 == *limit {
			break
		}

		control_number := record[0]
		lastname := record[2]
		firstname := record[3]
		middlename := record[4]
		archive_location := record[5]

		student := entities.NewStudent(
			control_number,
			firstname,
			middlename,
			lastname, "",
			entities.NON_TRANSFEREE,
			entities.SINGLE,
			"",
			"",
			archive_location,
		)

		if _, err := s.studentRepo.Create(ctx, student); err != nil {
			return err
		}

		log.Println("inserted: ", student.FullName())
	}

	return nil
}

func (s *StudentSeeder) readFile() ([][]string, error) {
	file, err := os.Open(s.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}
