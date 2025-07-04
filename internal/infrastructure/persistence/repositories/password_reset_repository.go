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

type PasswordResetRepository struct {
	db *sqlx.DB
}

func NewPasswordResetRepository(db *sqlx.DB) *PasswordResetRepository {
	return &PasswordResetRepository{
		db: db,
	}
}

func (r *PasswordResetRepository) Create(ctx context.Context, request *entities.PasswordResetRequest) (*entities.PasswordResetRequest, error) {
	query, args, err := sq.Insert("reset_request").
		Columns("id", "admin_id", "expires_at", "status", "created_at", "updated_at").
		Values(request.ID, request.Admin.ID, request.ExpiresAt, request.Status, request.CreatedAt, request.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return request, nil
}

func (r *PasswordResetRepository) FindByID(ctx context.Context, id string) (*entities.PasswordResetRequest, error) {
	query, args, err := sq.Select("*").
		From("reset_request").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	request := new(models.PasswordResetRequestModel)
	if err := r.db.GetContext(ctx, request, query, args...); err != nil {
		return nil, err
	}

	{
		query, args, err := sq.Select("*").
			From("admin").
			Where(sq.Eq{"id": request.AdminID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		if err := r.db.GetContext(ctx, &request.Admin, query, args...); err != nil {
			return nil, err
		}
	}

	return request.ToDomain(), nil
}

func (r *PasswordResetRepository) FindByAdminIDWhereActive(ctx context.Context, adminID string) (*entities.PasswordResetRequest, error) {
	query, args, err := sq.Select("*").
		From("reset_request").
		Where(
			sq.And{
				sq.Eq{"admin_id": adminID},
				sq.Eq{"status": entities.REQUEST_STATUS_PENDING},
			},
		).
		ToSql()
	if err != nil {
		return nil, err
	}

	request := new(models.PasswordResetRequestModel)
	if err := r.db.GetContext(ctx, request, query, args...); err != nil {
		return nil, err
	}

	{
		query, args, err := sq.Select("*").
			From("admin").
			Where(sq.Eq{"id": request.AdminID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		if err := r.db.GetContext(ctx, &request.Admin, query, args...); err != nil {
			return nil, err
		}
	}

	return request.ToDomain(), nil
}

func (r *PasswordResetRepository) FindAll(ctx context.Context) ([]*entities.PasswordResetRequest, error) {
	query, args, err := sq.Select("*").
		From("reset_request").
		Where(sq.Eq{"status": entities.REQUEST_STATUS_PENDING}).
		OrderBy("created_at ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	requests := make([]*models.PasswordResetRequestModel, 0)
	if err := r.db.SelectContext(ctx, &requests, query, args...); err != nil {
		return nil, err
	}

	for _, request := range requests {
		query, args, err := sq.Select("*").
			From("admin").
			Where(sq.Eq{"id": request.AdminID}).
			ToSql()
		if err != nil {
			return nil, err
		}

		if err := r.db.GetContext(ctx, &request.Admin, query, args...); err != nil {
			return nil, err
		}
	}

	results := slices.AppendSeq(
		make([]*entities.PasswordResetRequest, 0),
		utils.Map(requests, func(request *models.PasswordResetRequestModel) *entities.PasswordResetRequest {
			return request.ToDomain()
		}),
	)

	return results, nil
}

func (r *PasswordResetRepository) Save(ctx context.Context, request *entities.PasswordResetRequest) error {
	query, args, err := sq.Update("reset_request").
		SetMap(map[string]interface{}{
			"expires_at": request.ExpiresAt,
			"status":     request.Status,
			"updated_at": request.UpdatedAt,
		}).
		Where(sq.Eq{"id": request.ID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
