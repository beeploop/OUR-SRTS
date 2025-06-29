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
		Columns("id", "name").
		Values(docType.ID, docType.Name).
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
