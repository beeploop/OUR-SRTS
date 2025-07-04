package viewmodel

import (
	"fmt"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type StudentListItem struct {
	ControlNumber string `json:"control_number"`
	Firstname     string `json:"firstname"`
	Middlename    string `json:"middlename"`
	Lastname      string `json:"lastname"`
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
	ControlNumber string    `json:"control_number"`
	FirstName     string    `json:"firstname"`
	MiddleName    string    `json:"middlename"`
	LastName      string    `json:"lastname"`
	Suffix        string    `json:"suffix"`
	StudentType   string    `json:"student_type"`
	CivilStatus   string    `json:"civil_status"`
	Program       string    `json:"program"`
	Major         string    `json:"major"`
	Envelope      Envelope  `json:"envelope"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
		Envelope:      EnvelopeFromDomain(&student.Envelope),
		CreatedAt:     student.CreatedAt,
		UpdatedAt:     student.UpdatedAt,
	}
}

func (s Student) Fullname() string {
	return fmt.Sprintf("%s %s %s", s.FirstName, s.MiddleName, s.LastName)
}
