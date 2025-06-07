package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Major struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMajor(name string) *Major {
	return &Major{
		ID:    uuid.New().String(),
		Title: name,
	}
}

func (m *Major) validate() error {
	if m.Title == "" {
		return errors.New("name must not be empty")
	}
	if m.CreatedAt.After(m.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}
	return nil
}

func (m *Major) GetID() string {
	return m.ID
}

func (m *Major) GetTitle() string {
	return m.Title
}

func (m *Major) UpdateTitle(title string) error {
	m.Title = title
	m.UpdatedAt = time.Now()

	return m.validate()
}

func (m *Major) Copy() *Major {
	copy := &Major{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return copy
}
