package models

import "github.com/beeploop/our-srts/internal/domain/entities"

type DocumentTypeModel struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (m *DocumentTypeModel) ToDomain() *entities.DocumentType {
	return &entities.DocumentType{
		ID:   m.ID,
		Name: m.Name,
	}
}
