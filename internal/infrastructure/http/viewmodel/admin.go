package viewmodel

import "github.com/beeploop/our-srts/internal/infrastructure/session"

type Admin struct {
	ID       string
	Fullname string
	Username string
	Role     string
}

func AdminModelFromSession(admin *session.SessionModel) Admin {
	return Admin{
		ID:       admin.ID,
		Fullname: admin.Fullname,
		Username: admin.Username,
		Role:     admin.Role,
	}
}
