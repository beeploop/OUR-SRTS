package repositories

import (
	"context"
	"slices"

	sq "github.com/Masterminds/squirrel"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/models"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) Create(ctx context.Context, student *entities.Student) (*entities.Student, error) {
	columns := []string{
		"control_number",
		"first_name",
		"middle_name",
		"last_name",
		"suffix",
		"student_type",
		"civil_status",
		"program_id",
		"major_id",
		"archive_location",
		"created_at",
		"updated_at",
	}

	values := []interface{}{
		student.ControlNumber,
		student.FirstName,
		student.MiddleName,
		student.LastName,
		student.Suffix,
		student.StudentType,
		student.CivilStatus,
		student.ProgramID,
		student.MajorID,
		student.ArchiveLocation,
		student.CreatedAt,
		student.UpdatedAt,
	}

	query, args, err := sq.Insert("student").
		Columns(columns...).
		Values(values...).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return student, nil
}

func (r *StudentRepository) Search(ctx context.Context, filter repositories.StudentFilter) ([]*entities.Student, error) {
	queryBuilder := sq.Select("*").From("student")

	switch filter.SearchType {
	case repositories.SEARCH_BY_FIRSTNAME:
		queryBuilder.Where(sq.Eq{"first_name": filter.Query})
	case repositories.SEARCH_BY_MIDDLENAME:
		queryBuilder.Where(sq.Eq{"middle_name": filter.Query})
	case repositories.SEARCH_BY_LASTNAME:
		queryBuilder.Where(sq.Eq{"last_name": filter.Query})
	case repositories.SEARCH_BY_CONTROL_NUMBER:
		queryBuilder.Where(sq.Eq{"control_number": filter.Query})
	}

	if filter.ProgramID != "" {
		queryBuilder.Where(sq.Eq{"program_id": filter.ProgramID})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	students := make([]*models.StudentModel, 0)
	if err := r.db.SelectContext(ctx, &students, query, args...); err != nil {
		return nil, err
	}

	results := slices.AppendSeq(
		make([]*entities.Student, 0),
		utils.Map(students, func(student *models.StudentModel) *entities.Student {
			return student.ToDomain()
		}),
	)

	return results, nil
}

func (r *StudentRepository) Save(ctx context.Context, student *entities.Student) error {
	query, args, err := sq.Update("student").
		SetMap(map[string]interface{}{
			"first_name":       student.FirstName,
			"middle_name":      student.MiddleName,
			"last_name":        student.LastName,
			"suffix":           student.Suffix,
			"student_type":     student.StudentType,
			"civil_status":     student.CivilStatus,
			"program_id":       student.ProgramID,
			"major_id":         student.MajorID,
			"archive_location": student.ArchiveLocation,
			"created_at":       student.CreatedAt,
			"updated_at":       student.UpdatedAt,
		}).
		Where(sq.Eq{"control_number": student.ControlNumber}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
