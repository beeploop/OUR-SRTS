package viewmodel

import "github.com/beeploop/our-srts/internal/domain/entities"

type Program struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func ProgramFromDomain(program *entities.Program) Program {
	return Program{
		ID:    program.ID,
		Title: program.Title,
	}
}

type Major struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func MajorFromDomain(major *entities.Major) Major {
	return Major{
		ID:    major.ID,
		Title: major.Title,
	}
}

type ProgramWithMajors struct {
	Program Program `json:"program"`
	Majors  []Major `json:"majors"`
}

func GetProgramWithTitle(programs []ProgramWithMajors, title string) Program {
	for _, program := range programs {
		if program.Program.Title == title {
			return program.Program
		}
	}

	return Program{}
}
