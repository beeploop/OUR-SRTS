package viewmodel

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type Document struct {
	ID          string
	Type        string
	Filename    string
	StoragePath string
	UploadedAt  time.Time
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
