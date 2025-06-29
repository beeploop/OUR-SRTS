package seeder

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type DocumentTypeSeeder struct {
	documentTypes    []entities.DocumentType
	documentTypeRepo repositories.DocumentTypeRepository
}

func NewDocumentTypeSeeder(documents []entities.DocumentType, repo repositories.DocumentTypeRepository) *DocumentTypeSeeder {
	return &DocumentTypeSeeder{
		documentTypes:    documents,
		documentTypeRepo: repo,
	}
}

func (s *DocumentTypeSeeder) Execute(ctx context.Context) error {
	for _, document := range s.documentTypes {
		if _, err := s.documentTypeRepo.Create(ctx, &document); err != nil {
			return err
		}
	}

	return nil
}
