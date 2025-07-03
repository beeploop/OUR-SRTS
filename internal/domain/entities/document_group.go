package entities

import (
	"slices"

	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type DocumentGroup struct {
	Type      DocumentType
	Documents []Document
}

func NewDocumentGroup(doctype DocumentType) *DocumentGroup {
	return &DocumentGroup{
		Type:      doctype,
		Documents: make([]Document, 0),
	}
}

func (g *DocumentGroup) Append(document Document) {
	g.Documents = append(g.Documents, document)
}

func (g *DocumentGroup) Copy() *DocumentGroup {
	return &DocumentGroup{
		Type: g.Type,
		Documents: slices.AppendSeq(
			make([]Document, 0),
			utils.Map(g.Documents, func(document Document) Document {
				return *document.Copy()
			}),
		),
	}
}
