package viewmodel

import (
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
)

type Admin struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Enabled  bool   `json:"enabled"`
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
