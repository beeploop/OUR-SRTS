package store

import "database/sql"

type Program struct {
	Program string   `db:"program"`
	Majors  []string `db:"major"`
}

func GetProgramsAndMajors() ([]Program, error) {
	query := `
        SELECT
          p.program AS program,
          m.major AS major
        FROM
          Program p
          LEFT JOIN Major m ON m.programId = p.id
        ORDER BY 
            p.program
    `

	programs := []Program{}

	rows, err := Db_Conn.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currentProgram := Program{}
	for rows.Next() {
		var programName string
		var majorName sql.NullString

		err := rows.Scan(&programName, &majorName)
		if err != nil {
			return nil, err
		}

		if currentProgram.Program != programName || len(currentProgram.Majors) == 0 {
			if currentProgram.Program != "" {
				programs = append(programs, currentProgram)
			}
			currentProgram = Program{
				Program: programName,
				Majors:  []string{},
			}
		}

		if majorName.Valid {
			currentProgram.Majors = append(currentProgram.Majors, majorName.String)
		}
	}

	if currentProgram.Program != "" {
		programs = append(programs, currentProgram)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}
