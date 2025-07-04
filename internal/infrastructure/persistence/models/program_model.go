package models

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type ProgramModel struct {
	ID        string       `db:"id"`
	Title     string       `db:"title"`
	Majors    []MajorModel `db:"majors"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

func (m *ProgramModel) ToDomain() *entities.Program {
	return &entities.Program{
		ID:    m.ID,
		Title: m.Title,
		Majors: slices.AppendSeq(
			make([]entities.Major, 0),
			utils.Map(m.Majors, func(major MajorModel) entities.Major {
				return *major.ToDomain()
			}),
		),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.CreatedAt,
	}
}
