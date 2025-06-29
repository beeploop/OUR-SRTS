package reset

import (
	"context"
	"errors"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
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

func (u *UseCase) FulfillRequest(ctx context.Context, requestID, newPassword, password string) error {
	session, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return errors.New("unauthorized access")
	}

	admin, err := u.adminRepo.FindById(ctx, session.ID)
	if err != nil {
		return err
	}

	if !admin.IsPasswordCorrect(password) {
		return errors.New("unauthorized access")
	}

	request, err := u.passwordResetRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	requestor, err := u.adminRepo.FindById(ctx, request.Admin.ID)
	if err != nil {
		return err
	}

	if err := requestor.UpdatePassword(newPassword); err != nil {
		return err
	}

	if err := request.Fulfill(); err != nil {
		return err
	}
	if err := u.passwordResetRepo.Save(ctx, request); err != nil {
		return err
	}

	if err := u.adminRepo.Save(ctx, requestor); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) RejectRequest(ctx context.Context, requestID, password string) error {
	session, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return errors.New("unauthorized access")
	}

	admin, err := u.adminRepo.FindById(ctx, session.ID)
	if err != nil {
		return err
	}

	if !admin.IsPasswordCorrect(password) {
		return errors.New("unauthorized access")
	}

	request, err := u.passwordResetRepo.FindByID(ctx, requestID)
	if err != nil {
		return err
	}

	if err := request.Reject(); err != nil {
		return err
	}

	if err := u.passwordResetRepo.Save(ctx, request); err != nil {
		return err
	}

	return nil
}
