package store

import (
	"strconv"

	"github.com/BeepLoop/registrar-digitized/types"
)

func SearchStudent(searchTerm, program string) ([]types.Student, error) {
	// check if search term can be converted to number in order to know
	// if we should search by lastname or by controlNumber
	startingChars := searchTerm[:2]
	_, err := strconv.Atoi(startingChars)
	if err != nil {
		// if program == all then dont filter by program
		// else filter students by program
		if program == "all" {
			students, err := getByLastname(searchTerm)
			if err != nil {
				return nil, err
			}
			return students, nil
		}

		// program != all so we doesn't filter students by program
		students, err := getByFilteredLastname(searchTerm, program)
		if err != nil {
			return nil, err
		}
		return students, nil
	}

	// if program == all then dont filter by program
	// else filter students by program
	if program == "all" {
		students, err := getByControlNo(searchTerm)
		if err != nil {
			return nil, err
		}
		return students, nil
	}

	// program != all so we filter by program
	students, err := getByFilteredControlNo(searchTerm, program)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func getByLastname(lastname string) ([]types.Student, error) {
	query := "select * from Student where lastname=?"

	var students []types.Student
	err := Db_Conn.Select(&students, query, lastname)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByFilteredLastname(lastname, program string) ([]types.Student, error) {
	query := `
        SELECT
            *
        FROM
            Student
        WHERE
            lastname = ?
            AND programId = (
                SELECT
                    id
                FROM
                    Program
                WHERE
                    program = ?
            );
    `

	students := []types.Student{}
	err := Db_Conn.Select(&students, query, lastname, program)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByControlNo(controlNumber string) ([]types.Student, error) {
	query := "select * from Student where controlNumber=?"

	var students []types.Student
	err := Db_Conn.Select(&students, query, controlNumber)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByFilteredControlNo(controlNumber, program string) ([]types.Student, error) {
	query := `
        SELECT
            *
        FROM
            Student
        WHERE
            controlNumber = ?
            AND programId = (
                SELECT
                    id
                FROM
                    Program
                WHERE
                    program = ?
            );
    `

	students := []types.Student{}
	err := Db_Conn.Select(&students, query, controlNumber, program)
	if err != nil {
		return nil, err
	}

	return students, nil
}
