package entities

import (
	"errors"
	"fmt"
	"time"
)

type StudentType string
type CivilStatus string

const (
	NON_TRANSFEREE StudentType = "non_transferee"
	TRANSFEREE     StudentType = "transferee"
	GRADUATE       StudentType = "graduate"

	SINGLE  CivilStatus = "single"
	MARRIED CivilStatus = "married"
)

type Student struct {
	ControlNumber string
	FirstName     string
	MiddleName    string
	LastName      string
	Suffix        string
	StudentType   StudentType
	CivilStatus   CivilStatus
	ProgramID     string
	MajorID       string
	Program       Program
	Major         Major
	ImagePath     string
	Envelope      Envelope
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewStudent(
	controlNumber string,
	firstname string,
	middlename string,
	lastname string,
	suffix string,
	studentType StudentType,
	status CivilStatus,
	programID string,
	majorID string,
	fileLocation string,
) *Student {
	return &Student{
		ControlNumber: controlNumber,
		FirstName:     firstname,
		MiddleName:    middlename,
		LastName:      lastname,
		Suffix:        suffix,
		StudentType:   studentType,
		CivilStatus:   status,
		ProgramID:     programID,
		MajorID:       majorID,
		Envelope:      *NewEnvelope(fmt.Sprintf("%s_%s", controlNumber, lastname), fileLocation),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (s *Student) Validate() error {
	if s.FirstName == "" {
		return errors.New("firstname must not be empty")
	}
	if s.LastName == "" {
		return errors.New("lastname must not be empty")
	}
	if s.ProgramID == "" {
		return errors.New("student must have a program")
	}
	if s.CreatedAt.After(s.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}

	return s.Envelope.Validate()
}

func (s *Student) FullName() string {
	if s.Suffix != "" {
		return fmt.Sprintf("%s %s %s, %s", s.FirstName, s.MiddleName, s.LastName, s.Suffix)
	}
	return fmt.Sprintf("%s %s %s", s.FirstName, s.MiddleName, s.LastName)
}

func (s *Student) FullUpdate(student *Student) error {
	s.FirstName = student.FirstName
	s.MiddleName = student.MiddleName
	s.LastName = student.LastName
	s.Suffix = student.Suffix
	s.StudentType = student.StudentType
	s.CivilStatus = student.CivilStatus
	s.ProgramID = student.ProgramID
	s.MajorID = student.MajorID
	s.Envelope.UpdateLocation(student.Envelope.Location)
	s.Envelope.UpdateOwner(fmt.Sprintf("%s_%s", student.ControlNumber, student.LastName))
	s.UpdatedAt = student.UpdatedAt
	return s.Validate()
}

func (s *Student) UpdateName(firstname, middlename, lastname, suffix string) error {
	s.FirstName = firstname
	s.MiddleName = middlename
	s.LastName = lastname
	s.Suffix = suffix
	return s.Validate()
}

func (s *Student) UpdateType(studentType StudentType) error {
	s.StudentType = studentType
	return s.Validate()
}

func (s *Student) UpdateCivilStatus(civilStatus CivilStatus) error {
	s.CivilStatus = civilStatus
	return s.Validate()
}

func (s *Student) AddDocument(document Document) {
	s.Envelope.AddDocument(document)
}

func (s *Student) Copy() *Student {
	return &Student{
		ControlNumber: s.ControlNumber,
		FirstName:     s.FirstName,
		MiddleName:    s.MiddleName,
		LastName:      s.LastName,
		Suffix:        s.Suffix,
		StudentType:   s.StudentType,
		CivilStatus:   s.CivilStatus,
		ProgramID:     s.ProgramID,
		MajorID:       s.MajorID,
		Envelope:      *s.Envelope.Copy(),
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
	}
}
