package viewmodel

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type Document struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Filename    string    `json:"filename"`
	StoragePath string    `json:"storage_path"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

func DocumentFromDomain(document *entities.Document) Document {
	return Document{
		ID:          document.ID,
		Type:        document.Type.Name,
		Filename:    document.Filename,
		StoragePath: document.StoragePath,
		UploadedAt:  document.UploadedAt,
	}
}
