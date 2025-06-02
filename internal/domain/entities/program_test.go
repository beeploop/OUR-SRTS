package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgram(t *testing.T) {
	t.Run("test update title", func(t *testing.T) {
		initial := "initial title"
		updated := "updated title"

		p := NewProgram(initial)
		err := p.UpdatTitle(updated)

		assert.NoError(t, err)
		assert.EqualValues(t, updated, p.Title)
	})

	t.Run("test add major", func(t *testing.T) {
		p := NewProgram("test program title")
		major := NewMajor("new major")

		err := p.AddMajor(*major)

		assert.NoError(t, err)
		assert.Contains(t, p.Majors, *major)
	})

	t.Run("test update major title", func(t *testing.T) {
		p := NewProgram("test program title")
		major := NewMajor("test major")
		input := "updated major title"

		err := p.AddMajor(*major)
		assert.NoError(t, err)

		err = p.UpdateMajorTitle(major.GetID(), input)
		assert.NoError(t, err)

		assert.EqualValues(t, input, p.Majors[0].GetTitle())
	})
}
