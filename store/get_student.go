package store

import (
	"github.com.BeepLoop/registrar-digitized/types"
)

func GetStudent(controlNumber string) (types.Student, error) {
	query := "SELECT * FROM Student WHERE controlNumber = ?"

	var student types.Student
	err := Db_Conn.Get(&student, query, controlNumber)
	if err != nil {
		return student, err
	}

	return student, nil
}
