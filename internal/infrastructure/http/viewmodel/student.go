package viewmodel

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type StudentListItem struct {
	ControlNumber string
	Firstname     string
	Middlename    string
	Lastname      string
}

func StudentItemFromDomain(student *entities.Student) StudentListItem {
	return StudentListItem{
		ControlNumber: student.ControlNumber,
		Firstname:     student.FirstName,
		Middlename:    student.MiddleName,
		Lastname:      student.LastName,
	}
}

type Student struct {
	ControlNumber string
	FirstName     string
	MiddleName    string
	LastName      string
	Suffix        string
	StudentType   string
	CivilStatus   string
	Program       string
	Major         string
	Envelope      Envelope
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func StudentFromDomain(student *entities.Student) Student {
	return Student{
		ControlNumber: student.ControlNumber,
		FirstName:     student.FirstName,
		MiddleName:    student.MiddleName,
		LastName:      student.LastName,
		Suffix:        student.Suffix,
		StudentType:   string(student.StudentType),
		CivilStatus:   string(student.CivilStatus),
		Program:       student.Program.Title,
		Major:         student.Major.Title,
		Envelope:      EnvelopeFromDomain(student.Envelope),
		CreatedAt:     student.CreatedAt,
		UpdatedAt:     student.UpdatedAt,
	}
}
