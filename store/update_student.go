package store

import "github.com/BeepLoop/registrar-digitized/types"

func UpdateStudent(student types.StudentInfo) error {
	query := `
        WITH cte_program AS (
            SELECT id FROM Program WHERE program = ?
        )
        UPDATE Student 
        SET 
            lastname = ?,
            firstname = ?,
            middlename = ?,
            type = ?,
            civilStatus = ?,
            fileLocation = ?,
            programId = ( SELECT id FROM cte_program ),
            majorId = ( SELECT id FROM Major WHERE major = ?  AND programId = ( SELECT id FROM cte_program ))
        WHERE controlNumber = ?
    `

	_, err := Db_Conn.Exec(
		query,
		student.Program,
		student.Lastname,
		student.Firstname,
		student.Middlename,
		student.Type,
		student.CivilStatus,
		student.FileLocaion,
		student.Major,
		student.ControlNumber,
	)
	if err != nil {
		return err
	}

	return nil
}
