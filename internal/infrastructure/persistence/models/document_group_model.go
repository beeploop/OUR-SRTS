package models

import (
	"slices"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type DocumentGroupModel struct {
	Type      DocumentTypeModel
	Documents []DocumentModel
}

func (d *DocumentGroupModel) ToDomain() *entities.DocumentGroup {
	return &entities.DocumentGroup{
		Type: *d.Type.ToDomain(),
		Documents: slices.AppendSeq(
			make([]entities.Document, 0),
			utils.Map(d.Documents, func(document DocumentModel) entities.Document {
				return *document.ToDomain()
			}),
		),
	}
}
