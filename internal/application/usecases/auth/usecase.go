package auth

import (
	"context"
	"errors"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type UseCase struct {
	adminRepo repositories.AdminRepository
}

func NewUseCase(adminRepo repositories.AdminRepository) *UseCase {
	return &UseCase{
		adminRepo: adminRepo,
	}
}

func (u *UseCase) Login(ctx context.Context, username, password string) (*entities.Admin, error) {
	admin, err := u.adminRepo.FindByUsername(ctx, username)
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
