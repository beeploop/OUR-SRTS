package reset

import (
	"context"
	"errors"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type UseCase struct {
	adminRepo         repositories.AdminRepository
	passwordResetRepo repositories.PasswordResetRepository
}

func NewUseCase(
	adminRepo repositories.AdminRepository,
	passwordResetRepo repositories.PasswordResetRepository,
) *UseCase {
	return &UseCase{
		adminRepo:         adminRepo,
		passwordResetRepo: passwordResetRepo,
	}
}

func (u *UseCase) RequestPasswordReset(ctx context.Context, username string) error {
	if username == "" {
		return errors.New("invalid username")
	}

	account, err := u.adminRepo.FindByUsername(ctx, username)
	if err != nil {
		return err
	}

	if !account.Enabled {
		return errors.New("forbidden action")
	}

	request := entities.NewResetRequest(*account)
	if _, err := u.passwordResetRepo.Create(ctx, request); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) GetRequestList(ctx context.Context) ([]*entities.PasswordResetRequest, error) {
	return u.passwordResetRepo.FindAll(ctx)
}
