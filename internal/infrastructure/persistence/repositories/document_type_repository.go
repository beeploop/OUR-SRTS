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

type DocumentTypeRepository struct {
	db *sqlx.DB
}

func NewDocumentTypeRepository(db *sqlx.DB) *DocumentTypeRepository {
	return &DocumentTypeRepository{
		db: db,
	}
}

func (r *DocumentTypeRepository) Create(ctx context.Context, docType *entities.DocumentType) (*entities.DocumentType, error) {
	query, args, err := sq.Insert("document_type").
		Columns("id", "title", "full_title", "is_stable", "allow_multiple").
		Values(docType.ID, docType.Title, docType.FullTitle, docType.Stable, docType.AllowMultiple).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return docType, nil
}

func (r *DocumentTypeRepository) FindAll(ctx context.Context) ([]*entities.DocumentType, error) {
	query, args, err := sq.Select("*").
		From("document_type").
		ToSql()
	if err != nil {
		return nil, err
	}

	documentTypes := make([]*models.DocumentTypeModel, 0)
	if err := r.db.SelectContext(ctx, &documentTypes, query, args...); err != nil {
		return nil, err
	}

	results := slices.AppendSeq(
		make([]*entities.DocumentType, 0),
		utils.Map(documentTypes, func(docType *models.DocumentTypeModel) *entities.DocumentType {
			return docType.ToDomain()
		}),
	)

	return results, nil
}

func (r *DocumentTypeRepository) FindByTitle(ctx context.Context, title string) (*entities.DocumentType, error) {
	query, args, err := sq.Select("*").
		From("document_type").
		Where(sq.Eq{"title": title}).
		ToSql()
	if err != nil {
		return nil, err
	}

	documentType := new(models.DocumentTypeModel)
	if err := r.db.GetContext(ctx, documentType, query, args...); err != nil {
		return nil, err
	}

	return documentType.ToDomain(), nil
}
