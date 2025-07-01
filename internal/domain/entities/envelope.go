package entities

import (
	"errors"
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/google/uuid"
)

type Envelope struct {
	ID        string
	Owner     string
	Location  string
	Documents []Document
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEnvelope(owner, location string) *Envelope {
	return &Envelope{
		ID:        uuid.New().String(),
		Owner:     owner,
		Location:  location,
		Documents: make([]Document, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (e *Envelope) Validate() error {
	if e.Owner == "" {
		return errors.New("owner must not be empty")
	}
	if e.Location == "" {
		return errors.New("file location must be set")
	}
	return nil
}

func (e *Envelope) AddDocument(document Document) {
	e.Documents = append(e.Documents, document)
	e.UpdatedAt = time.Now()
}

func (e *Envelope) UpdateLocation(location string) error {
	if location == "" {
		return errors.New("invalid location")
	}

	e.Location = location
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Envelope) Copy() *Envelope {
	return &Envelope{
		ID:       e.ID,
		Owner:    e.Owner,
		Location: e.Location,
		Documents: slices.AppendSeq(
			make([]Document, 0),
			utils.Map(e.Documents, func(document Document) Document {
				return document
			}),
		),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
