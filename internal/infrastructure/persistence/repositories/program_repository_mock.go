package repositories

import (
	"context"
	"database/sql"
	"slices"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type ProgramRepositoryMock struct {
	programs []entities.Program
}

func NewProgramRepositoryMock() *ProgramRepositoryMock {
	return &ProgramRepositoryMock{
		programs: make([]entities.Program, 0),
	}
}

func (r *ProgramRepositoryMock) Create(ctx context.Context, program *entities.Program) (*entities.Program, error) {
	r.programs = append(r.programs, *program)
	return program, nil
}

func (r *ProgramRepositoryMock) FindById(ctx context.Context, id string) (*entities.Program, error) {
	for _, program := range r.programs {
		if program.ID == id {
			return &program, nil
		}
	}

	return nil, sql.ErrNoRows
}

func (r *ProgramRepositoryMock) FindAll(ctx context.Context) ([]*entities.Program, error) {
	programs := slices.AppendSeq(
		make([]*entities.Program, 0),
		utils.Map(r.programs, func(program entities.Program) *entities.Program {
			return &program
		}),
	)

	return programs, nil
}

func (r *ProgramRepositoryMock) Save(ctx context.Context, program *entities.Program) error {
	for i := range r.programs {
		if r.programs[i].ID == program.ID {
			r.programs[i] = *program
			return nil
		}
	}

	return nil
}
