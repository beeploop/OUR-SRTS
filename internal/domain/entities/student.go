package entities

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/pkg/utils"
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
	ControlNumber   string
	FirstName       string
	MiddleName      string
	LastName        string
	Suffix          string
	StudentType     StudentType
	CivilStatus     CivilStatus
	ProgramID       string
	MajorID         string
	ArchiveLocation string
	Documents       []Document
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
	archiveLocation string,
) *Student {
	return &Student{
		ControlNumber:   controlNumber,
		FirstName:       firstname,
		MiddleName:      middlename,
		LastName:        lastname,
		Suffix:          suffix,
		StudentType:     studentType,
		CivilStatus:     status,
		ProgramID:       programID,
		MajorID:         majorID,
		ArchiveLocation: archiveLocation,
		Documents:       make([]Document, 0),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
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

	return nil
}

func (s *Student) FullName() string {
	if s.Suffix != "" {
		return fmt.Sprintf("%s %s %s, %s", s.FirstName, s.MiddleName, s.LastName, s.Suffix)
	}
	return fmt.Sprintf("%s %s %s", s.FirstName, s.MiddleName, s.LastName)
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
	s.Documents = append(s.Documents, document)
}

func (s *Student) Copy() *Student {
	return &Student{
		ControlNumber:   s.ControlNumber,
		FirstName:       s.FirstName,
		MiddleName:      s.MiddleName,
		LastName:        s.LastName,
		Suffix:          s.Suffix,
		StudentType:     s.StudentType,
		CivilStatus:     s.CivilStatus,
		ProgramID:       s.ProgramID,
		MajorID:         s.MajorID,
		ArchiveLocation: s.ArchiveLocation,
		Documents: slices.AppendSeq(
			make([]Document, 0),
			utils.Map(s.Documents, func(document Document) Document {
				return document
			}),
		),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
