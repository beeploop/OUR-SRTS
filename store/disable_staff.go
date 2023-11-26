package store

func DisableStaff(username string) error {

	query := `
        UPDATE 
            User 
        SET 
            status = 'disabled'
        WHERE 
            username = ?
    `

	_, err := Db_Conn.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}
