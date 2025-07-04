package entities

import (
	"errors"
	"slices"
	"time"

	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/google/uuid"
)

type Envelope struct {
	ID             string
	Owner          string
	Location       string
	DocumentGroups []*DocumentGroup
	CreatedAt      time.Time
	updatedAt      time.Time
}

func NewEnvelope(owner, location string) *Envelope {
	return &Envelope{
		ID:             uuid.New().String(),
		Owner:          owner,
		Location:       location,
		DocumentGroups: make([]*DocumentGroup, 0),
		CreatedAt:      time.Now(),
		updatedAt:      time.Now(),
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
	for _, group := range e.DocumentGroups {
		if group.Type.Title == document.Type.Title {
			group.Append(document)
			return
		}
	}

	group := NewDocumentGroup(document.Type)
	group.Append(document)
	e.DocumentGroups = append(e.DocumentGroups, group)
}

func (e *Envelope) UpdateOwner(owner string) error {
	e.Owner = owner
	e.updatedAt = time.Now()
	return e.Validate()
}

func (e *Envelope) UpdateLocation(location string) error {
	if location == "" {
		return errors.New("invalid location")
	}

	e.Location = location
	e.updatedAt = time.Now()
	return e.Validate()
}

func (e *Envelope) GroupWithTypeTitle(title string) *DocumentGroup {
	for _, group := range e.DocumentGroups {
		if group.Type.Title == title {
			return group
		}
	}

	return nil
}

func (s *Envelope) SetUpdatedAt(t time.Time) error {
	s.updatedAt = t
	return s.Validate()
}

func (s *Envelope) UpdatedAt() time.Time {
	return s.updatedAt
}

func (e *Envelope) Copy() *Envelope {
	return &Envelope{
		ID:       e.ID,
		Owner:    e.Owner,
		Location: e.Location,
		DocumentGroups: slices.AppendSeq(
			make([]*DocumentGroup, 0),
			utils.Map(e.DocumentGroups, func(group *DocumentGroup) *DocumentGroup {
				return group.Copy()
			}),
		),
		CreatedAt: e.CreatedAt,
		updatedAt: e.updatedAt,
	}
}
