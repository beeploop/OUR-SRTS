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

type Major struct {
	ID    string
	Title string
}

func MajorFromDomain(major *entities.Major) Major {
	return Major{
		ID:    major.ID,
		Title: major.Title,
	}
}

type ProgramWithMajors struct {
	Program Program
	Majors  []Major
}
