package seeder

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type DocumentTypeModel struct {
	Title         string `json:"title"`
	FullTitle     string `json:"full_title"`
	Stable        bool   `json:"is_stable"`
	AllowMultiple bool   `json:"allow_multiple"`
}

type DocumentTypes struct {
	Documents []DocumentTypeModel `json:"documents"`
}

type DocumentTypeSeeder struct {
	sourceFile       string
	documentTypeRepo repositories.DocumentTypeRepository
}

func NewDocumentTypeSeeder(sourceFile string, repo repositories.DocumentTypeRepository) *DocumentTypeSeeder {
	return &DocumentTypeSeeder{
		sourceFile:       sourceFile,
		documentTypeRepo: repo,
	}
}

func (s *DocumentTypeSeeder) Execute(ctx context.Context) error {
	b, err := s.readFile()
	if err != nil {
		return err
	}

	documentTypes := new(DocumentTypes)
	if err := json.Unmarshal(b, documentTypes); err != nil {
		return err
	}

	for _, t := range documentTypes.Documents {
		document := entities.NewDocumentType(t.Title, t.FullTitle, t.Stable, t.AllowMultiple)

		if _, err := s.documentTypeRepo.Create(ctx, document); err != nil {
			return err
		}
	}

	return nil
}

func (s *DocumentTypeSeeder) readFile() ([]byte, error) {
	file, err := os.Open(s.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
