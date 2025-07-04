package repositories

import (
	"context"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type SearchType string

const (
	SEARCH_BY_FIRSTNAME      SearchType = "firstname"
	SEARCH_BY_MIDDLENAME     SearchType = "middlename"
	SEARCH_BY_LASTNAME       SearchType = "lastname"
	SEARCH_BY_CONTROL_NUMBER SearchType = "control_number"
)

type StudentFilter struct {
	Query      string
	SearchType SearchType
	ProgramID  string
}

type StudentRepository interface {
	Create(ctx context.Context, student *entities.Student) (*entities.Student, error)
	UploadDocument(ctx context.Context, document *entities.Document, envelope *entities.Envelope) (*entities.Document, error)
	FindByControlNumber(ctx context.Context, controlNumber string) (*entities.Student, error)
	Search(ctx context.Context, filter StudentFilter) ([]*entities.Student, error)
	Save(ctx context.Context, student *entities.Student) error
}
