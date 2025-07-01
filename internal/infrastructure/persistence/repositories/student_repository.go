package repositories

import (
	"context"
	"database/sql"
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
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	{
		// Create envelope
		envelope := student.Envelope
		query, args, err := sq.Insert("envelope").
			Columns("id", "owner", "location", "created_at", "updated_at").
			Values(envelope.ID, envelope.Owner, envelope.Location, envelope.CreatedAt, envelope.UpdatedAt).
			ToSql()
		if err != nil {
			return nil, err
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return nil, err
		}
	}

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
		"envelope_id",
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
		student.Envelope.ID,
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

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return student, nil
}

func (r *StudentRepository) FindByControlNumber(ctx context.Context, controlNumber string) (*entities.Student, error) {
	query, args, err := sq.Select("*").
		From("student").
		Where(sq.Eq{"control_number": controlNumber}).
		ToSql()
	if err != nil {
		return nil, err
	}

	student := new(models.StudentModel)
	if err := r.db.GetContext(ctx, student, query, args...); err != nil {
		return nil, err
	}

	{
		// Retrieve program and major
		query1, args1, err := sq.Select("*").
			From("program").
			Where(sq.Eq{"id": student.ProgramID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		if err := r.db.GetContext(ctx, &student.Program, query1, args1...); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}

		query2, args2, err := sq.Select("*").
			From("major").
			Where(sq.Eq{"id": student.MajorID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		if err := r.db.GetContext(ctx, &student.Major, query2, args2...); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}
	}

	if envelope, err := r.findEnvelope(ctx, student.EnvelopeID); err != nil {
		return nil, err
	} else {
		student.Envelope = *envelope
	}

	return student.ToDomain(), nil
}

func (r *StudentRepository) Search(ctx context.Context, filter repositories.StudentFilter) ([]*entities.Student, error) {
	queryBuilder := sq.Select("*").From("student")

	switch filter.SearchType {
	case repositories.SEARCH_BY_FIRSTNAME:
		queryBuilder = queryBuilder.Where(sq.Like{"first_name": "%" + filter.Query + "%"})
	case repositories.SEARCH_BY_MIDDLENAME:
		queryBuilder = queryBuilder.Where(sq.Like{"middle_name": "%" + filter.Query + "%"})
	case repositories.SEARCH_BY_LASTNAME:
		queryBuilder = queryBuilder.Where(sq.Like{"last_name": "%" + filter.Query + "%"})
	case repositories.SEARCH_BY_CONTROL_NUMBER:
		queryBuilder = queryBuilder.Where(sq.Like{"control_number": "%" + filter.Query + "%"})
	}

	if filter.ProgramID != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"program_id": filter.ProgramID})
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
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query1, args1, err := sq.Update("envelope").
		SetMap(map[string]interface{}{
			"owner":      student.Envelope.Owner,
			"location":   student.Envelope.Location,
			"updated_at": student.Envelope.UpdatedAt,
		}).
		Where(sq.Eq{"id": student.Envelope.ID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, query1, args1...); err != nil {
		return err
	}

	query2, args2, err := sq.Update("student").
		SetMap(map[string]interface{}{
			"first_name":   student.FirstName,
			"middle_name":  student.MiddleName,
			"last_name":    student.LastName,
			"suffix":       student.Suffix,
			"student_type": student.StudentType,
			"civil_status": student.CivilStatus,
			"program_id":   student.ProgramID,
			"major_id":     student.MajorID,
			"created_at":   student.CreatedAt,
			"updated_at":   student.UpdatedAt,
		}).
		Where(sq.Eq{"control_number": student.ControlNumber}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, query2, args2...); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *StudentRepository) findEnvelope(ctx context.Context, envelopeID string) (*models.EnvelopeModel, error) {
	query, args, err := sq.Select("*").
		From("envelope").
		Where(sq.Eq{"id": envelopeID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	envelope := new(models.EnvelopeModel)
	if err := r.db.GetContext(ctx, envelope, query, args...); err != nil {
		return nil, err
	}

	documentTypes, err := r.getDocumentTypes(ctx)
	if err != nil {
		return nil, err
	}

	{
		for _, docType := range documentTypes {
			query, args, err := sq.Select("*").
				From("document").
				Where(sq.And{
					sq.Eq{"envelope_id": envelopeID},
					sq.Eq{"type_id": docType.ID},
				}).
				ToSql()
			if err != nil {
				return nil, err
			}

			document := new(models.DocumentModel)
			if err := r.db.GetContext(ctx, document, query, args...); err != nil {
				if err != sql.ErrNoRows {
					return nil, err
				}
			}

			envelope.Documents = append(envelope.Documents, models.DocumentModel{
				ID:          document.ID,
				TypeID:      docType.ID,
				Type:        *docType,
				EnvelopeID:  envelopeID,
				Filename:    document.Filename,
				StoragePath: document.StoragePath,
				UploadedAt:  document.UploadedAt,
			})
		}
	}

	return envelope, nil
}

func (r *StudentRepository) getDocumentTypes(ctx context.Context) ([]*models.DocumentTypeModel, error) {
	query, args, err := sq.Select("*").
		From("document_type").
		OrderBy("name ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	documentTypes := make([]*models.DocumentTypeModel, 0)
	if err := r.db.SelectContext(ctx, &documentTypes, query, args...); err != nil {
		return nil, err
	}

	return documentTypes, nil
}
