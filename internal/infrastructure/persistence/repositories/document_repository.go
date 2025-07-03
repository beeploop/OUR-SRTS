package repositories

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence/models"
	"github.com/jmoiron/sqlx"
)

type DocumentRepository struct {
	db *sqlx.DB
}

func NewDocumentRepository(db *sqlx.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

func (r *DocumentRepository) FindByID(ctx context.Context, id string) (*entities.Document, error) {
	query, args, err := sq.Select("*").
		From("document").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	document := new(models.DocumentModel)
	if err := r.db.GetContext(ctx, document, query, args...); err != nil {
		return nil, err
	}

	return document.ToDomain(), nil
}

func (r *DocumentRepository) Save(ctx context.Context, document *entities.Document) error {
	query, args, err := sq.Update("document").
		SetMap(map[string]interface{}{
			"filename":     document.Filename,
			"storage_path": document.StoragePath,
			"uploaded_at":  document.UploadedAt,
		}).
		Where(sq.Eq{"id": document.ID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
