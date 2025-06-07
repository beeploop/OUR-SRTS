package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AdminRole string

const (
	ROLE_SUPER_ADMIN AdminRole = "super_admin"
	ROLE_STAFF       AdminRole = "staff"
)

type Admin struct {
	ID        string
	Fullname  string
	Username  string
	Password  string
	Role      AdminRole
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAdmin(fullname, username, password string, role AdminRole) *Admin {
	return &Admin{
		ID:        uuid.New().String(),
		Fullname:  fullname,
		Username:  username,
		Password:  password,
		Role:      role,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Admin) validate() error {
	if a.Fullname == "" {
		return errors.New("fullname must not be empty")
	}
	if a.Username == "" {
		return errors.New("username must not be empty")
	}
	if a.Password == "" {
		return errors.New("password must not be empty")
	}
	return nil
}

func (a *Admin) UpdateName(fullname string) error {
	a.Fullname = fullname
	a.UpdatedAt = time.Now()
	return a.validate()
}

func (a *Admin) UpdateUsername(username string) error {
	a.Username = username
	a.UpdatedAt = time.Now()
	return a.validate()
}

func (a *Admin) UpdatePassword(password string) error {
	a.Password = password
	a.UpdatedAt = time.Now()
	return a.validate()
}

func (a *Admin) Disable() error {
	a.Enabled = false
	a.UpdatedAt = time.Now()
	return a.validate()
}

func (a *Admin) Enable() error {
	a.Enabled = true
	a.UpdatedAt = time.Now()
	return a.validate()
}
