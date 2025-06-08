package viewmodel

import "github.com/beeploop/our-srts/internal/domain/entities"

type StudentListItem struct {
	ID         string
	Firstname  string
	Middlename string
	Lastname   string
}

func StudentItemFromDomain(student *entities.Student) StudentListItem {
	return StudentListItem{
		ID:         student.ID,
		Firstname:  student.FirstName,
		Middlename: student.MiddleName,
		Lastname:   student.LastName,
	}
}
