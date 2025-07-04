package entities

import "github.com/google/uuid"

type DocumentType struct {
	ID            string
	Title         string
	FullTitle     string
	Stable        bool
	AllowMultiple bool
}

func NewDocumentType(title, fullTitle string, stable, allowMultiple bool) *DocumentType {
	return &DocumentType{
		ID:            uuid.New().String(),
		Title:         title,
		FullTitle:     fullTitle,
		Stable:        stable,
		AllowMultiple: allowMultiple,
	}
}

func (d *DocumentType) Copy() *DocumentType {
	return &DocumentType{
		ID:            d.ID,
		Title:         d.Title,
		FullTitle:     d.FullTitle,
		Stable:        d.Stable,
		AllowMultiple: d.AllowMultiple,
	}
}
