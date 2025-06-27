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

func (u *UseCase) CreateStaff(ctx context.Context, admin *entities.Admin) error {
	if err := admin.Validate(); err != nil {
		return err
	}

	if _, err := u.adminRepo.Create(ctx, admin); err != nil {
		return err
	}

	return nil
}
