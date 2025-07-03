package models

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type StudentModel struct {
	ControlNumber string               `db:"control_number"`
	FirstName     string               `db:"first_name"`
	MiddleName    string               `db:"middle_name"`
	LastName      string               `db:"last_name"`
	Suffix        string               `db:"suffix"`
	StudentType   entities.StudentType `db:"student_type"`
	CivilStatus   entities.CivilStatus `db:"civil_status"`
	ProgramID     string               `db:"program_id"`
	Program       ProgramModel         `db:"program"`
	MajorID       string               `db:"major_id"`
	Major         MajorModel           `db:"major"`
	EnvelopeID    string               `db:"envelope_id"`
	Envelope      EnvelopeModel        `db:"envelope"`
	Documents     []DocumentModel      `db:"documents"`
	CreatedAt     time.Time            `db:"created_at"`
	UpdatedAt     time.Time            `db:"updated_at"`
}

func (m *StudentModel) ToDomain() *entities.Student {
	return &entities.Student{
		ControlNumber: m.ControlNumber,
		FirstName:     m.FirstName,
		MiddleName:    m.MiddleName,
		LastName:      m.LastName,
		Suffix:        m.Suffix,
		StudentType:   m.StudentType,
		CivilStatus:   m.CivilStatus,
		ProgramID:     m.ProgramID,
		Program:       *m.Program.ToDomain(),
		MajorID:       m.MajorID,
		Major:         *m.Major.ToDomain(),
		Envelope:      *m.Envelope.ToDomain(),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}
