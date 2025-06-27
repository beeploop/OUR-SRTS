package viewmodel

import (
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
)

type Admin struct {
	ID       string
	Fullname string
	Username string
	Role     string
	Enabled  bool
}

func AdminModelFromSession(admin *session.SessionModel) Admin {
	return Admin{
		ID:       admin.ID,
		Fullname: admin.Fullname,
		Username: admin.Username,
		Role:     admin.Role,
		Enabled:  admin.Enabled,
	}
}

func AdminFromDomain(admin *entities.Admin) Admin {
	return Admin{
		ID:       admin.ID,
		Fullname: admin.Fullname,
		Username: admin.Username,
		Role:     string(admin.Role),
		Enabled:  admin.Enabled,
	}
}
