package seeder

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
)

type AdminModel struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Admins struct {
	Admins []AdminModel `json:"admins"`
}

type AdminSeeder struct {
	sourceFile string
	adminRepo  repositories.AdminRepository
}

func NewAdminSeeder(sourceFile string, repo repositories.AdminRepository) *AdminSeeder {
	return &AdminSeeder{
		sourceFile: sourceFile,
		adminRepo:  repo,
	}
}

func (s *AdminSeeder) Execute(ctx context.Context) error {
	b, err := s.readFile()
	if err != nil {
		return err
	}

	admins := new(Admins)
	if err := json.Unmarshal(b, admins); err != nil {
		return err
	}

	for _, admin := range admins.Admins {
		a := entities.NewAdmin(admin.Fullname, admin.Username, admin.Password, entities.AdminRole(admin.Role))

		if _, err := s.adminRepo.Create(ctx, a); err != nil {
			return err
		}
	}

	return nil
}

func (s *AdminSeeder) readFile() ([]byte, error) {
	file, err := os.Open(s.sourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
