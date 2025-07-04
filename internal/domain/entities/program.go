package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Program struct {
	ID        string
	Title     string
	Majors    []Major
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProgram(title string) *Program {
	return &Program{
		ID:        uuid.New().String(),
		Title:     title,
		Majors:    make([]Major, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Program) validate() error {
	if p.Title == "" {
		return errors.New("name must not be empty")
	}
	if p.CreatedAt.After(p.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}
	return nil
}

func (p *Program) UpdatTitle(title string) error {
	p.Title = title
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Program) AddMajor(major Major) error {
	p.Majors = append(p.Majors, major)
	p.UpdatedAt = time.Now()

	return p.validate()
}

func (p *Program) UpdateMajorTitle(majorID, title string) error {
	for i, major := range p.Majors {
		if major.GetID() == majorID {
			copy := *major.Copy()
			copy.UpdateTitle(title)

			p.Majors[i] = copy
			return nil
		}
	}

	return fmt.Errorf("major with id: %s, not found in program: %s", majorID, p.ID)
}
