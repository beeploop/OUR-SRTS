package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type AdminRepository interface {
	Create(ctx context.Context, admin *entities.Admin) (*entities.Admin, error)
	FindById(ctx context.Context, id string) (*entities.Admin, error)
	FindByUsername(ctx context.Context, username string) (*entities.Admin, error)
	FindAll(ctx context.Context) ([]*entities.Admin, error)
	Save(ctx context.Context, admin *entities.Admin) error
	Delete(ctx context.Context, id string) error
}
