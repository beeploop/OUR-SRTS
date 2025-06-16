package student

import (
	"context"
	"net/url"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type UseCase struct {
	studentRepo repositories.StudentRepository
}

func NewUseCase(studentRepo repositories.StudentRepository) *UseCase {
	return &UseCase{
		studentRepo: studentRepo,
	}
}

func (u *UseCase) AddStudent(ctx context.Context, student *entities.Student) error {
	if _, err := u.studentRepo.Create(ctx, student); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) Search(ctx context.Context, params url.Values) ([]*entities.Student, error) {
	if !params.Has("query") || params.Get("query") == "" {
		return make([]*entities.Student, 0), nil
	}

	if !params.Has("program") {
		params.Set("program", "all")
	}

	if !params.Has("type") {
		params.Set("type", string(repositories.SEARCH_BY_FIRSTNAME))
	}

	filter := repositories.StudentFilter{
		Query:      params.Get("query"),
		SearchType: repositories.SearchType(params.Get("type")),
		ProgramID:  utils.Ternary(params.Get("program") != "all", params.Get("program"), ""),
	}

	return u.studentRepo.Search(ctx, filter)
}
