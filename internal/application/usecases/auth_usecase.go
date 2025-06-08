package usecases

import (
	"errors"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type AuthUseCase struct {
	adminRepo repositories.AdminRepository
}

func NewAuthUseCase(adminRepo repositories.AdminRepository) *AuthUseCase {
	return &AuthUseCase{
		adminRepo: adminRepo,
	}
}

func (u *AuthUseCase) Login(username, password string) (*entities.Admin, error) {
	admin, err := u.adminRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if !admin.IsPasswordCorrect(password) {
		return nil, errors.New("invalid password")
	}

	if !admin.Enabled {
		return nil, errors.New("account disabled")
	}

	return admin, nil
}
