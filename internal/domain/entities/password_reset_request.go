package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ResetRequestStatus string

const (
	REQUEST_STATUS_PENDING   ResetRequestStatus = "pending"
	REQUEST_STATUS_FULFILLED ResetRequestStatus = "fulfilled"
	REQUEST_STATUS_REJECTED  ResetRequestStatus = "rejected"
)

type PasswordResetRequest struct {
	ID        string
	AdminID   string
	Token     string
	ExpiresAt time.Time
	Status    ResetRequestStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewResetRequest(adminID, token string, expiresAt time.Time) *PasswordResetRequest {
	return &PasswordResetRequest{
		ID:        uuid.New().String(),
		AdminID:   adminID,
		Token:     token,
		ExpiresAt: expiresAt,
		Status:    REQUEST_STATUS_PENDING,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *PasswordResetRequest) validate() error {
	if r.AdminID == "" {
		return errors.New("admin ID must not be empty")
	}
	if r.Token == "" {
		return errors.New("token must not be empty")
	}
	if r.ExpiresAt.Before(time.Now()) {
		return errors.New("expires_at must not be before current time")
	}
	return nil
}

func (r *PasswordResetRequest) Fulfill() error {
	if r.Status == REQUEST_STATUS_REJECTED {
		return errors.New("rejected requests cannot be fulfilled")
	}

	if r.Status == REQUEST_STATUS_FULFILLED {
		return nil
	}

	r.Status = REQUEST_STATUS_FULFILLED
	r.UpdatedAt = time.Now()
	return r.validate()
}

func (r *PasswordResetRequest) Reject() error {
	if r.Status == REQUEST_STATUS_FULFILLED {
		return errors.New("fulfilled requests cannot be rejected")
	}

	if r.Status == REQUEST_STATUS_REJECTED {
		return nil
	}

	r.Status = REQUEST_STATUS_REJECTED
	r.UpdatedAt = time.Now()
	return r.validate()
}

func (r *PasswordResetRequest) IsExpired() bool {
	return r.ExpiresAt.Before(time.Now()) || r.ExpiresAt.Equal(time.Now())
}
