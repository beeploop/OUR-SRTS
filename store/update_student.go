package store

import "github.com/registrar/types"

func UpdateStudent(student types.StudentInfo) error {
	query := `
        UPDATE Student 
        SET 
            lastname = ?,
            firstname = ?,
            middlename = ?,
            type = ?,
            civilStatus = ?,
            fileLocation = ?,
            programId = ( SELECT id FROM Program WHERE program = ? ),
            majorId = ( SELECT id FROM Major WHERE major = ? )
        WHERE controlNumber = ?
    `

	_, err := Db_Conn.Exec(
		query,
		student.Lastname,
		student.Firstname,
		student.Middlename,
		student.Type,
		student.CivilStatus,
		student.FileLocaion,
		student.Program,
		student.Major,
		student.ControlNumber,
	)
    if err != nil {
        return err
    }

	return nil
}
