package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          string
	Type        DocumentType
	Filename    string
	StoragePath string
	UploadedAt  time.Time
}

func NewDocument(docType DocumentType, filename, storagePath string) *Document {
	return &Document{
		ID:          uuid.New().String(),
		Type:        docType,
		Filename:    filename,
		StoragePath: storagePath,
		UploadedAt:  time.Now(),
	}
}

func (d *Document) validate() error {
	if d.Filename == "" {
		return errors.New("filename must be provided")
	}
	if d.StoragePath == "" {
		return errors.New("storage path must be provided")
	}
	return nil
}

func (d *Document) UpdateStoragePath(storagePath string) error {
	d.StoragePath = storagePath
	return d.validate()
}

func (d *Document) UpdateFilename(filename string) error {
	d.Filename = filename
	return d.validate()
}
