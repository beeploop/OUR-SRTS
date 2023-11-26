package store

func FulfillRequest(requestId, newPassword string) error {
	type Requestor struct {
		RequestorId string `db:"requestorId"`
	}

	query1 := `
        UPDATE
            Request 
        SET status = 'fulfilled'
        WHERE id = ?;
    `

	query2 := `
        SELECT
            r.requestorId as requestorId
        FROM Request r
        WHERE r.id = ?;
    `

	tx, err := Db_Conn.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query1, requestId)
	if err != nil {
		tx.Rollback()
		return err
	}

	var requestor Requestor
	err = tx.Get(&requestor, query2, requestId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE User SET password = ? WHERE id = ?", newPassword, requestor.RequestorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
