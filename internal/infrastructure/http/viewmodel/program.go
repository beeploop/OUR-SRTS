package viewmodel

import "github.com/beeploop/our-srts/internal/domain/entities"

type Program struct {
	ID    string
	Title string
}

func ProgramFromDomain(program *entities.Program) Program {
	return Program{
		ID:    program.ID,
		Title: program.Title,
	}
}
