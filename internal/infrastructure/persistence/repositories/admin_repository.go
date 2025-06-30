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

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) Create(ctx context.Context, admin *entities.Admin) (*entities.Admin, error) {
	query, args, err := sq.Insert("admin").
		Columns("id", "fullname", "username", "password", "role", "enabled", "created_at", "updated_at").
		Values(admin.ID, admin.Fullname, admin.Username, admin.Password, admin.Role, admin.Enabled, admin.CreatedAt, admin.UpdatedAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepository) FindById(ctx context.Context, id string) (*entities.Admin, error) {
	query, args, err := sq.Select("*").
		From("admin").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	admin := new(models.AdminModel)
	if err := r.db.GetContext(ctx, admin, query, args...); err != nil {
		return nil, err
	}

	return admin.ToDomain(), nil
}

func (r *AdminRepository) FindByUsername(ctx context.Context, username string) (*entities.Admin, error) {
	query, args, err := sq.Select("*").
		From("admin").
		Where(sq.Eq{"username": username}).
		ToSql()
	if err != nil {
		return nil, err
	}

	admin := new(models.AdminModel)
	if err := r.db.GetContext(ctx, admin, query, args...); err != nil {
		return nil, err
	}

	return admin.ToDomain(), nil
}

func (r *AdminRepository) FindAll(ctx context.Context) ([]*entities.Admin, error) {
	query, args, err := sq.Select("*").
		From("admin").
		OrderBy("role ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	modelAccounts := make([]*models.AdminModel, 0)
	if err := r.db.SelectContext(ctx, &modelAccounts, query, args...); err != nil {
		return nil, err
	}

	domainAccounts := slices.AppendSeq(
		make([]*entities.Admin, 0),
		utils.Map(modelAccounts, func(account *models.AdminModel) *entities.Admin {
			return account.ToDomain()
		}),
	)

	return domainAccounts, nil
}

func (r *AdminRepository) Save(ctx context.Context, admin *entities.Admin) error {
	query, args, err := sq.Update("admin").
		SetMap(map[string]interface{}{
			"fullname":   admin.Fullname,
			"username":   admin.Username,
			"password":   admin.Password,
			"role":       admin.Role,
			"enabled":    admin.Enabled,
			"created_at": admin.CreatedAt,
			"updated_at": admin.UpdatedAt,
		}).
		Where(sq.Eq{"id": admin.ID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) Delete(ctx context.Context, id string) error {
	query, args, err := sq.Delete("admin").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
