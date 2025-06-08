package repositories

import (
	"github.com/beeploop/our-srts/internal/domain/entities"
)

type AdminRepositoryMock struct{}

func NewAdminRepositoryMock() *AdminRepositoryMock {
	return &AdminRepositoryMock{}
}

func (r *AdminRepositoryMock) Create(admin *entities.Admin) (*entities.Admin, error) {
	return admin, nil
}

func (r *AdminRepositoryMock) FindById(id string) (*entities.Admin, error) {
	admin := entities.NewAdmin("firstname lastname", "admin", "password", entities.ROLE_SUPER_ADMIN)
	return admin, nil
}

func (r *AdminRepositoryMock) FindByUsername(username string) (*entities.Admin, error) {
	admin := entities.NewAdmin("firstname lastname", username, "password", entities.ROLE_SUPER_ADMIN)
	return admin, nil
}

func (r *AdminRepositoryMock) Save(admin *entities.Admin) error {
	return nil
}
