package store

import "github.com/lithammer/shortuuid/v4"

func NewRequest(username string) error {
	uuid := shortuuid.New()

	query := `
        INSERT INTO 
            Request (id, requestorId) 
        VALUES (
            ?,
            (SELECT id FROM User WHERE username = ?)
        )
    `

	_, err := Db_Conn.Exec(query, uuid, username)
	if err != nil {
		return err
	}

	return nil
}
