package repositories

import "github.com/beeploop/our-srts/internal/domain/entities"

type AdminRepository interface {
	Create(admin *entities.Admin) (*entities.Admin, error)
	FindById(id string) (*entities.Admin, error)
	FindByUsername(username string) (*entities.Admin, error)
	Save(admin *entities.Admin) error
}
