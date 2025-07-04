package viewmodel

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type PasswordResetRequest struct {
	ID        string    `json:"id"`
	Admin     Admin     `json:"admin"`
	Status    string    `json:"status"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
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
