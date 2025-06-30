package viewmodel

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type Envelope struct {
	ID        string
	Owner     string
	Location  string
	Documents []Document
	CreatedAt time.Time
	UpdatedAt time.Time
}

func EnvelopeFromDomain(envelope *entities.Envelope) Envelope {
	return Envelope{
		ID:       envelope.ID,
		Owner:    envelope.Owner,
		Location: envelope.Location,
		Documents: slices.AppendSeq(
			make([]Document, 0),
			utils.Map(envelope.Documents, func(document entities.Document) Document {
				return DocumentFromDomain(&document)
			}),
		),
		CreatedAt: envelope.CreatedAt,
		UpdatedAt: envelope.UpdatedAt,
	}
}

func (e Envelope) DocumentWithType(docType string) Document {
	for _, document := range e.Documents {
		if document.Type == docType {
			return document
		}

		break
	}

	return Document{}
}
