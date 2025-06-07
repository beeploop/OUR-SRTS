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
	id        string
	adminId   string
	token     string
	expiresAt time.Time
	status    ResetRequestStatus
	createdAt time.Time
	updatedAt time.Time
}

func NewResetRequest(adminID, token string, expiresAt time.Time) *PasswordResetRequest {
	return &PasswordResetRequest{
		id:        uuid.New().String(),
		adminId:   adminID,
		token:     token,
		expiresAt: expiresAt,
		status:    REQUEST_STATUS_PENDING,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (r *PasswordResetRequest) validate() error {
	if r.adminId == "" {
		return errors.New("admin ID must not be empty")
	}
	if r.token == "" {
		return errors.New("token must not be empty")
	}
	if r.expiresAt.Before(time.Now()) {
		return errors.New("expires_at must not be before current time")
	}
	return nil
}

func (r *PasswordResetRequest) ID() string {
	return r.id
}

func (r *PasswordResetRequest) AdminID() string {
	return r.adminId
}

func (r *PasswordResetRequest) Token() string {
	return r.token
}

func (r *PasswordResetRequest) Status() ResetRequestStatus {
	return r.status
}

func (r *PasswordResetRequest) Fulfill() error {
	if r.status == REQUEST_STATUS_REJECTED {
		return errors.New("rejected requests cannot be fulfilled")
	}

	if r.status == REQUEST_STATUS_FULFILLED {
		return nil
	}

	r.status = REQUEST_STATUS_FULFILLED
	r.updatedAt = time.Now()
	return r.validate()
}

func (r *PasswordResetRequest) Reject() error {
	if r.status == REQUEST_STATUS_FULFILLED {
		return errors.New("fulfilled requests cannot be rejected")
	}

	if r.status == REQUEST_STATUS_REJECTED {
		return nil
	}

	r.status = REQUEST_STATUS_REJECTED
	r.updatedAt = time.Now()
	return r.validate()
}

func (r *PasswordResetRequest) IsExpired() bool {
	return r.expiresAt.Before(time.Now()) || r.expiresAt.Equal(time.Now())
}
