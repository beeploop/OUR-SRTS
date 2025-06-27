package admin

import (
	"context"

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

func (u *UseCase) GetAccounts(ctx context.Context) ([]*entities.Admin, error) {
	return u.adminRepo.FindAll(ctx)
}
