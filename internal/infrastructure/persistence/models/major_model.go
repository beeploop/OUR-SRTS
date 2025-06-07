package models

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type MajorModel struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *MajorModel) ToDomain() *entities.Major {
	return &entities.Major{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
