package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Major struct {
	id        string
	title     string
	createdAt time.Time
	updatedAt time.Time
}

func NewMajor(name string) *Major {
	return &Major{
		id:    uuid.New().String(),
		title: name,
	}
}

func (m *Major) validate() error {
	if m.title == "" {
		return errors.New("name must not be empty")
	}
	if m.createdAt.After(m.updatedAt) {
		return errors.New("created_at must be before updated_at")
	}
	return nil
}

func (m *Major) GetID() string {
	return m.id
}

func (m *Major) GetTitle() string {
	return m.title
}

func (m *Major) UpdateTitle(title string) error {
	m.title = title
	m.updatedAt = time.Now()

	return m.validate()
}

func (m *Major) Copy() *Major {
	copy := &Major{
		id:        m.id,
		title:     m.title,
		createdAt: m.createdAt,
		updatedAt: m.updatedAt,
	}

	return copy
}
