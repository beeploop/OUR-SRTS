package store

func RejectRequest(requestId string) error {

	query := `
        UPDATE
            Request 
        SET status = 'rejected'
        WHERE id = ?

    `

	_, err := Db_Conn.Exec(query, requestId)
	if err != nil {
		return err
	}

	return nil
}
