package repositories

import (
	"context"
	"slices"

	sq "github.com/Masterminds/squirrel"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/models"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type ProgramRepository struct {
	db *sqlx.DB
}

func NewProgramRepository(db *sqlx.DB) *ProgramRepository {
	return &ProgramRepository{
		db: db,
	}
}

func (r *ProgramRepository) Create(ctx context.Context, program *entities.Program) (*entities.Program, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := sq.Insert("program").
		Columns("id", "title", "created_at", "updated_at").
		Values(program.ID, program.Title, program.CreatedAt, program.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	for _, major := range program.Majors {
		query, args, err := sq.Insert("major").
			Columns("id", "title", "program_id", "created_at", "updated_at").
			Values(major.ID, major.Title, program.ID, major.CreatedAt, major.UpdatedAt).
			ToSql()
		if err != nil {
			return nil, err
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return program, nil
}

func (r *ProgramRepository) FindById(ctx context.Context, id string) (*entities.Program, error) {
	query, args, err := sq.Select("*").
		From("program").
		Where(sq.Eq{"control_number": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	program := new(models.ProgramModel)

	if err := r.db.GetContext(ctx, program, query, args...); err != nil {
		return nil, err
	}

	return program.ToDomain(), nil
}

func (r *ProgramRepository) FindAll(ctx context.Context) ([]*entities.Program, error) {
	query, args, err := sq.Select("*").
		From("program").
		ToSql()
	if err != nil {
		return nil, err
	}

	programs := make([]*models.ProgramModel, 0)
	if err := r.db.SelectContext(ctx, &programs, query, args...); err != nil {
		return nil, err
	}

	results := slices.AppendSeq(
		make([]*entities.Program, 0),
		utils.Map(programs, func(program *models.ProgramModel) *entities.Program {
			return program.ToDomain()
		}),
	)

	return results, nil
}

func (r *ProgramRepository) Save(ctx context.Context, program *entities.Program) error {
	query, args, err := sq.Update("program").
		SetMap(map[string]interface{}{}).
		Where(sq.Eq{"id": program.ID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
