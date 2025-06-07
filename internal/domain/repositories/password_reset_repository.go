package repositories

import "github.com/beeploop/our-srts/internal/domain/entities"

type PasswordResetRepository interface {
	FindAll() ([]*entities.PasswordResetRequest, error)
	Save(request *entities.PasswordResetRequest) error
}
