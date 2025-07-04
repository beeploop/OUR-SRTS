package models

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type AdminModel struct {
	ID        string             `db:"id"`
	Fullname  string             `db:"fullname"`
	Username  string             `db:"username"`
	Password  string             `db:"password"`
	Role      entities.AdminRole `db:"role"`
	Enabled   bool               `db:"enabled"`
	CreatedAt time.Time          `db:"created_at"`
	UpdatedAt time.Time          `db:"updated_at"`
}

func (m *AdminModel) ToDomain() *entities.Admin {
	return &entities.Admin{
		ID:        m.ID,
		Fullname:  m.Fullname,
		Username:  m.Username,
		Password:  m.Password,
		Role:      m.Role,
		Enabled:   m.Enabled,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
