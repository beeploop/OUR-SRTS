package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStudent(t *testing.T) {
	testStudent := NewStudent(
		"id_foo",
		"jane",
		"john",
		"doe",
		"",
		NON_TRANSFEREE,
		SINGLE,
		"id_program",
		"id_major",
		"shelf_1",
	)

	t.Run("test get fullname", func(t *testing.T) {
		tests := []struct {
			input    *Student
			expected string
		}{
			{
				input:    NewStudent("id_foo", "foo", "bar", "baz", "", NON_TRANSFEREE, SINGLE, "program_foo", "major_foo", "shelf_1"),
				expected: "foo bar baz",
			},
			{
				input:    NewStudent("id_foo", "foo", "bar", "baz", "jr.", NON_TRANSFEREE, SINGLE, "program_foo", "major_foo", "shelf_1"),
				expected: "foo bar baz, jr.",
			},
			{
				input:    NewStudent("id_foo", "foo", "bar", "baz", "III", NON_TRANSFEREE, SINGLE, "program_foo", "major_foo", "shelf_1"),
				expected: "foo bar baz, III",
			},
		}

		for _, test := range tests {
			fullname := test.input.FullName()
			assert.EqualValues(t, test.expected, fullname)
		}
	})

	t.Run("test update name", func(t *testing.T) {
		tests := []struct {
			initial         *Student
			inputFirstname  string
			inputMiddlename string
			inputLastname   string
			inputSuffix     string
			expected        string
		}{
			{
				initial:         testStudent.Copy(),
				inputFirstname:  "juan",
				inputMiddlename: "dela",
				inputLastname:   "cruz",
				inputSuffix:     "",
				expected:        "juan dela cruz",
			},
			{
				initial:         testStudent.Copy(),
				inputFirstname:  "juan",
				inputMiddlename: "dela",
				inputLastname:   "cruz",
				inputSuffix:     "jr.",
				expected:        "juan dela cruz, jr.",
			},
			{
				initial:         testStudent.Copy(),
				inputFirstname:  "juan",
				inputMiddlename: "dela",
				inputLastname:   "cruz",
				inputSuffix:     "III",
				expected:        "juan dela cruz, III",
			},
		}

		for _, test := range tests {
			firstname := test.inputFirstname
			middlename := test.inputMiddlename
			lastname := test.inputLastname
			suffix := test.inputSuffix
			err := test.initial.UpdateName(firstname, middlename, lastname, suffix)

			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, test.initial.FullName())
		}
	})

	t.Run("test add document", func(t *testing.T) {
		student := testStudent.Copy()
		docT := DocumentType{
			ID:    "test",
			Title: "test",
		}
		document := NewDocument(docT, "test file", "/file/path")

		student.AddDocument(*document)

		group := student.Envelope.GroupWithTypeTitle("test")

		assert.NotNil(t, group)
		assert.EqualValues(t, 1, len(group.Documents))

		addedDoc := group.Documents[0]

		assert.ObjectsAreEqualValues(document, addedDoc)
	})

	t.Run("test full update", func(t *testing.T) {
		student := testStudent.Copy()
		updatedStudent := testStudent.Copy()
		updatedStudent.FirstName = "foo"
		updatedStudent.StudentType = GRADUATE
		updatedStudent.CivilStatus = MARRIED
		updatedStudent.UpdatedAt = time.Now()

		err := student.FullUpdate(updatedStudent)

		assert.NoError(t, err)
		assert.EqualExportedValues(t, student, updatedStudent)
	})
}
