package entities

import "github.com/google/uuid"

type DocumentType struct {
	ID   string
	Name string
}

func NewDocumentType(name string) *DocumentType {
	return &DocumentType{
		ID:   uuid.New().String(),
		Name: name,
	}
}
