package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type ProgramRepository interface {
	Create(ctx context.Context, program *entities.Program) (*entities.Program, error)
	FindById(ctx context.Context, id string) (*entities.Program, error)
	FindAll(ctx context.Context) ([]*entities.Program, error)
	Save(ctx context.Context, program *entities.Program) error
}
