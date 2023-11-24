package store

import "github.com/registrar/types"

func AddStudent(student types.StudentInfo) error {
	query := `
    INSERT INTO
    Student (
        controlNumber,
        lastname,
        firstname,
        middlename,
        type,
        civilStatus,
        fileLocation,
        programId,
        majorId
    )
    VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        (
            SELECT
                p.id
            FROM Program p
            WHERE p.program = ?
        ),
        (
            SELECT
                m.id
            FROM Major m
            WHERE m.major = ?
        )
    )
    `

	_, err := Db_Conn.Exec(
		query,
		student.ControlNumber,
		student.Lastname,
		student.Firstname,
		student.Middlename,
		student.Type,
		student.CivilStatus,
		student.FileLocaion,
		student.Program,
		student.Major,
	)

	if err != nil {
		return err
	}

	return nil
}
