package models

import (
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
)

type DocumentModel struct {
	ID          string            `db:"id"`
	TypeID      string            `db:"type_id"`
	Type        DocumentTypeModel `db:"document_type"`
	EnvelopeID  string            `db:"envelope_id"`
	Filename    string            `db:"filename"`
	StoragePath string            `db:"storage_path"`
	UploadedAt  time.Time         `db:"uploaded_at"`
}

func (m *DocumentModel) ToDomain() *entities.Document {
	return &entities.Document{
		ID:          m.ID,
		Type:        *m.Type.ToDomain(),
		Filename:    m.Filename,
		StoragePath: m.StoragePath,
		UploadedAt:  m.UploadedAt,
	}
}
