package admin

import (
	"context"
	"errors"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
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

func (u *UseCase) CreateAccount(ctx context.Context, admin *entities.Admin) error {
	if err := admin.Validate(); err != nil {
		return err
	}

	if _, err := u.adminRepo.Create(ctx, admin); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) DeleteAccount(ctx context.Context, accountID, password string) error {
	if accountID == "" {
		return errors.New("invalid account ID")
	}

	session, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return errors.New("unauthorized access")
	}

	admin, err := u.adminRepo.FindByUsername(ctx, session.Username)
	if err != nil {
		return err
	}

	if !admin.IsPasswordCorrect(password) {
		return errors.New("unauthorized access")
	}

	return u.adminRepo.Delete(ctx, accountID)
}
