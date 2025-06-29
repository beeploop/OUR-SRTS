package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, request *entities.PasswordResetRequest) (*entities.PasswordResetRequest, error)
	FindAll(ctx context.Context) ([]*entities.PasswordResetRequest, error)
	Save(ctx context.Context, request *entities.PasswordResetRequest) error
}
