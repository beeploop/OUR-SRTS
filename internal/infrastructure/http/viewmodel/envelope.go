package viewmodel

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type Envelope struct {
	ID             string          `json:"id"`
	Owner          string          `json:"owner"`
	Location       string          `json:"location"`
	DocumentGroups []DocumentGroup `json:"document_groups"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

func EnvelopeFromDomain(envelope *entities.Envelope) Envelope {
	return Envelope{
		ID:       envelope.ID,
		Owner:    envelope.Owner,
		Location: envelope.Location,
		DocumentGroups: slices.AppendSeq(
			make([]DocumentGroup, 0),
			utils.Map(envelope.DocumentGroups, func(group *entities.DocumentGroup) DocumentGroup {
				return DocumentGroupFromDomain(group)
			}),
		),
		CreatedAt: envelope.CreatedAt,
		UpdatedAt: envelope.UpdatedAt(),
	}
}

func (e Envelope) DocumentWithType(docType string) []Document {
	for _, group := range e.DocumentGroups {
		if group.Type.Title == docType {
			return group.Documents
		}
	}

	return make([]Document, 0)
}

func (e Envelope) SingleDocumentWithType(docType string) Document {
	for _, group := range e.DocumentGroups {
		if group.Type.Title == docType {
			if !group.IsEmpty() {
				return group.Documents[0]
			}
			break
		}
	}

	return Document{}
}
