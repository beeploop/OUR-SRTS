package store

import (
	"errors"

	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/sirupsen/logrus"
)

func SearchStudent(data types.SearchData) ([]types.SearchResult, error) {
	switch data.Type {
	case "lastname":
		if data.Program == "all" {
			logrus.Info("search by lastname, no program")
			return getByLastname(data.SearchTerm)
		} else {
			logrus.Info("search by lastname, with program")
			return getByFilteredLastname(data.SearchTerm, data.Program)
		}
	case "firstname":
		if data.Program == "all" {
			logrus.Info("search by firstname, no program")
			return getByFirstname(data.SearchTerm)
		} else {
			logrus.Info("search by firstname, with program")
			return getByFilteredFirstname(data.SearchTerm, data.Program)
		}
	default:
		return nil, errors.New("invalid search type")
	}
}

func getByLastname(lastname string) ([]types.SearchResult, error) {
	query := `
        SELECT 
            controlNumber, lastname, firstname, middlename 
        FROM 
            Student 
        WHERE 
            lastname LIKE ? 
        ORDER BY 
            lastname
    `

	var students []types.SearchResult
	err := Db_Conn.Select(&students, query, lastname+"%")
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByFilteredLastname(lastname, program string) ([]types.SearchResult, error) {
	query := `
        SELECT
            controlNumber, lastname, firstname, middlename 
        FROM
            Student
        WHERE
            lastname LIKE ?
            AND programId = (
                SELECT
                    id
                FROM
                    Program
                WHERE
                    program = ?
            )
        ORDER BY lastname
    `

	students := []types.SearchResult{}
	err := Db_Conn.Select(&students, query, lastname+"%", program)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByFirstname(firstname string) ([]types.SearchResult, error) {
	query := `
        SELECT 
            controlNumber, lastname, firstname, middlename 
        FROM 
            Student 
        WHERE 
            firstname LIKE ? 
        ORDER BY 
            lastname
    `

	var students []types.SearchResult
	err := Db_Conn.Select(&students, query, firstname+"%")
	if err != nil {
		return nil, err
	}

	return students, nil
}

func getByFilteredFirstname(firstname, program string) ([]types.SearchResult, error) {
	query := `
        SELECT
            controlNumber, lastname, firstname, middlename
        FROM
            Student
        WHERE
            firstname LIKE ?
            AND programId = (
                SELECT
                    id
                FROM
                    Program
                WHERE
                    program = ?
            )
        ORDER BY lastname
    `

	students := []types.SearchResult{}
	err := Db_Conn.Select(&students, query, firstname+"%", program)
	if err != nil {
		return nil, err
	}

	return students, nil
}
