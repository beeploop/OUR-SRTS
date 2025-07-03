package viewmodel

import "github.com/beeploop/our-srts/internal/domain/entities"

type DocumentType struct {
	Title         string `json:"title"`
	FullTitle     string `json:"full_title"`
	Stable        bool   `json:"is_stable"`
	AllowMultiple bool   `json:"allow_multiple"`
}

func DocumentTypeFromDomain(docType *entities.DocumentType) DocumentType {
	return DocumentType{
		Title:         docType.Title,
		FullTitle:     docType.FullTitle,
		Stable:        docType.Stable,
		AllowMultiple: docType.AllowMultiple,
	}
}
