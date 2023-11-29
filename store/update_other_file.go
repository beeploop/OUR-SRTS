package store

func UpdateOtherFile(fileLocation, filename string) error {

	query := `
        UPDATE 
            Other 
        SET 
            filename = ?,
            location = ?
        WHERE
            filename = ?
    `

	_, err := Db_Conn.Exec(query, filename, fileLocation, filename)
	if err != nil {
		return err
	}

	return nil
}

