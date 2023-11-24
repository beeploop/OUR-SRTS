package store

import (
	"fmt"

	"github.com/registrar/types"
)

func GetInitialStudents() []types.Student {
	query := "select * from Student limit 10"

	var students []types.Student

	err := Db_Conn.Select(&students, query)
	if err != nil {
		fmt.Println("error: ", err)
		panic(err)
	}

	fmt.Println("students: ", students)

	return students
}
