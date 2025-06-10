package seeder

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"slices"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type ProgramModel struct {
	Program string   `json:"program"`
	Majors  []string `json:"majors"`
}

type Programs struct {
	Programs []ProgramModel `json:"programs"`
}

type ProgramSeeder struct {
	sourceFile  string
	programRepo repositories.ProgramRepository
}

func NewProgramSeeder(sourceFile string, repo repositories.ProgramRepository) *ProgramSeeder {
	return &ProgramSeeder{
		sourceFile:  sourceFile,
		programRepo: repo,
	}
}

func (s *ProgramSeeder) Execute(ctx context.Context) error {
	b, err := s.readFile()
	if err != nil {
		return err
	}

	programs := new(Programs)
	if err := json.Unmarshal(b, programs); err != nil {
		return err
	}

	for _, program := range programs.Programs {
		p := entities.NewProgram(program.Program)

		majors := slices.AppendSeq(
			make([]entities.Major, 0),
			utils.Map(program.Majors, func(major string) entities.Major {
				return *entities.NewMajor(major)
			}),
		)

		p.Majors = majors

		if _, err := s.programRepo.Create(ctx, p); err != nil {
			return err
		}
	}

	return nil
}

func (s *ProgramSeeder) readFile() ([]byte, error) {
	file, err := os.Open(s.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
