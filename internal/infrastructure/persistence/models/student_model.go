package models

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type StudentModel struct {
	ID              string               `db:"id"`
	FirstName       string               `db:"first_name"`
	MiddleName      string               `db:"middle_name"`
	LastName        string               `db:"last_name"`
	Suffix          string               `db:"suffix"`
	StudentType     entities.StudentType `db:"student_type"`
	CivilStatus     entities.CivilStatus `db:"civil_status"`
	ProgramID       string               `db:"program_id"`
	MajorID         string               `db:"major_id"`
	ArchiveLocation string               `db:"archive_location"`
	Documents       []DocumentModel      `db:"documents"`
	CreatedAt       time.Time            `db:"created_at"`
	UpdatedAt       time.Time            `db:"updated_at"`
}

func (m *StudentModel) ToDomain() *entities.Student {
	return &entities.Student{
		ID:              m.ID,
		FirstName:       m.FirstName,
		MiddleName:      m.MiddleName,
		LastName:        m.LastName,
		Suffix:          m.Suffix,
		StudentType:     m.StudentType,
		CivilStatus:     m.CivilStatus,
		ProgramID:       m.ProgramID,
		MajorID:         m.MajorID,
		ArchiveLocation: m.ArchiveLocation,
		Documents: slices.AppendSeq(
			make([]entities.Document, 0),
			utils.Map(m.Documents, func(document DocumentModel) entities.Document {
				return *document.ToDomain()
			}),
		),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
