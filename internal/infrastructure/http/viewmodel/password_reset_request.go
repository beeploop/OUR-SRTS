package viewmodel

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type PasswordResetRequest struct {
	ID        string
	Admin     Admin
	Status    string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func PasswordResetRequestFromDomain(request *entities.PasswordResetRequest) PasswordResetRequest {
	return PasswordResetRequest{
		ID:        request.ID,
		Admin:     AdminFromDomain(&request.Admin),
		Status:    string(request.Status),
		ExpiresAt: request.ExpiresAt,
		CreatedAt: request.CreatedAt,
	}
}
