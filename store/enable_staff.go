package store

func EnableStaff(username string) error {

	query := `
        UPDATE 
            User 
        SET 
            status = 'enabled'
        WHERE 
            username = ?
    `

	_, err := Db_Conn.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}
