package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type DocumentRepository interface {
	FindByID(ctx context.Context, id string) (*entities.Document, error)
	Save(ctx context.Context, document *entities.Document) error
}
