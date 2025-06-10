package program

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type UseCase struct {
	programRepo repositories.ProgramRepository
}

func NewUseCase(programRepo repositories.ProgramRepository) *UseCase {
	return &UseCase{
		programRepo: programRepo,
	}
}

func (u *UseCase) GetProgramList(ctx context.Context) ([]*entities.Program, error) {
	return u.programRepo.FindAll(ctx)
}
