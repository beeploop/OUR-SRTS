package models

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type EnvelopeModel struct {
	ID        string          `db:"id"`
	Owner     string          `db:"owner"`
	Location  string          `db:"location"`
	Documents []DocumentModel `db:"documents"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}

func (e *EnvelopeModel) ToDomain() *entities.Envelope {
	return &entities.Envelope{
		ID:       e.ID,
		Owner:    e.Owner,
		Location: e.Location,
		Documents: slices.AppendSeq(
			make([]entities.Document, 0),
			utils.Map(e.Documents, func(document DocumentModel) entities.Document {
				return *document.ToDomain()
			}),
		),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
