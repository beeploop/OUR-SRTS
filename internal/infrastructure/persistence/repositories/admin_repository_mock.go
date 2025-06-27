package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type AdminRepositoryMock struct{}

func NewAdminRepositoryMock() *AdminRepositoryMock {
	return &AdminRepositoryMock{}
}

func (r *AdminRepositoryMock) Create(ctx context.Context, admin *entities.Admin) (*entities.Admin, error) {
	return admin, nil
}

func (r *AdminRepositoryMock) FindById(ctx context.Context, id string) (*entities.Admin, error) {
	admin := entities.NewAdmin("firstname lastname", "admin", "password", entities.ROLE_SUPER_ADMIN)
	return admin, nil
}

func (r *AdminRepositoryMock) FindByUsername(ctx context.Context, username string) (*entities.Admin, error) {
	admin := entities.NewAdmin("firstname lastname", username, "password", entities.ROLE_SUPER_ADMIN)
	return admin, nil
}

func (r *AdminRepositoryMock) FindAll(ctx context.Context) ([]*entities.Admin, error) {
	return make([]*entities.Admin, 0), nil
}

func (r *AdminRepositoryMock) Save(ctx context.Context, admin *entities.Admin) error {
	return nil
}
