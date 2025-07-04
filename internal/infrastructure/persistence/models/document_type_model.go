package models

import "github.com/beeploop/our-srts/internal/domain/entities"

type DocumentTypeModel struct {
	ID            string `db:"id"`
	Title         string `db:"title"`
	FullTitle     string `db:"full_title"`
	Stable        bool   `db:"is_stable"`
	AllowMultiple bool   `db:"allow_multiple"`
}

func (m *DocumentTypeModel) ToDomain() *entities.DocumentType {
	return &entities.DocumentType{
		ID:            m.ID,
		Title:         m.Title,
		FullTitle:     m.FullTitle,
		Stable:        m.Stable,
		AllowMultiple: m.AllowMultiple,
	}
}
