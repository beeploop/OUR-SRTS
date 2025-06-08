package models

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type PasswordResetRequestModel struct {
	ID        string                      `db:"id"`
	AdminID   string                      `db:"admin_id"`
	Token     string                      `db:"token"`
	ExpiresAt time.Time                   `db:"expires_at"`
	Status    entities.ResetRequestStatus `db:"status"`
	CreatedAt time.Time                   `db:"created_at"`
	UpdatedAt time.Time                   `db:"updated_at"`
}

func (m *PasswordResetRequestModel) ToDomain() *entities.PasswordResetRequest {
	return &entities.PasswordResetRequest{
		ID:        m.ID,
		AdminID:   m.AdminID,
		Token:     m.Token,
		ExpiresAt: m.ExpiresAt,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
