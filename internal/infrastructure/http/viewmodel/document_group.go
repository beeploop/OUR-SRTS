package viewmodel

import (
	"slices"

	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type DocumentGroup struct {
	Type      DocumentType `json:"type"`
	Documents []Document   `json:"documents"`
}

func DocumentGroupFromDomain(group *entities.DocumentGroup) DocumentGroup {
	return DocumentGroup{
		Type: DocumentTypeFromDomain(&group.Type),
		Documents: slices.AppendSeq(
			make([]Document, 0),
			utils.Map(group.Documents, func(document entities.Document) Document {
				return DocumentFromDomain(&document)
			}),
		),
	}
}

func (d *DocumentGroup) IsEmpty() bool {
	return len(d.Documents) == 0
}

func (d *DocumentGroup) First() Document {
	return d.Documents[0]
}
