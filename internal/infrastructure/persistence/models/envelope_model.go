package models

import (
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type EnvelopeModel struct {
	ID             string `db:"id"`
	Owner          string `db:"owner"`
	Location       string `db:"location"`
	DocumentGroups []DocumentGroupModel
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (e *EnvelopeModel) ToDomain() *entities.Envelope {
	envelope := &entities.Envelope{
		ID:       e.ID,
		Owner:    e.Owner,
		Location: e.Location,
		DocumentGroups: slices.AppendSeq(
			make([]*entities.DocumentGroup, 0),
			utils.Map(e.DocumentGroups, func(group DocumentGroupModel) *entities.DocumentGroup {
				return group.ToDomain()
			}),
		),
		CreatedAt: e.CreatedAt,
	}

	envelope.SetUpdatedAt(e.UpdatedAt)

	return envelope
}
