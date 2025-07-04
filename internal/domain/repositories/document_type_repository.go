package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type DocumentTypeRepository interface {
	Create(ctx context.Context, docType *entities.DocumentType) (*entities.DocumentType, error)
	FindAll(ctx context.Context) ([]*entities.DocumentType, error)
	FindByTitle(ctx context.Context, title string) (*entities.DocumentType, error)
}
