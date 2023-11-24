package store

func GetPrograms() ([]string, error) {

	query := `
        SELECT
            program 
        FROM 
            Program
    `

	programs := []string{}
	err := Db_Conn.Select(&programs, query)
	if err != nil {
		return nil, err
	}

	return programs, nil
}
