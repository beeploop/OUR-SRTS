package repositories

import "github.com/beeploop/our-srts/internal/domain/entities"

type SearchCriteria struct {
	FirstName  *string
	MiddleName *string
	LastName   *string
	ProgramID  *string
}

type StudentRepository interface {
	Create(student *entities.Student) (*entities.Student, error)
	Search(criteria SearchCriteria) ([]*entities.Student, error)
	FindAll(limit, offset int) ([]*entities.Student, error)
	Save(student *entities.Student) error
}
