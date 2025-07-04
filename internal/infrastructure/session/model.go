package session

import "github.com/beeploop/our-srts/internal/domain/entities"

type SessionModel struct {
	ID       string
	Fullname string
	Username string
	Role     string
	Enabled  bool
}

func FromDomain(admin *entities.Admin) SessionModel {
	return SessionModel{
		ID:       admin.ID,
		Fullname: admin.Fullname,
		Username: admin.Username,
		Role:     string(admin.Role),
		Enabled:  admin.Enabled,
	}
}
