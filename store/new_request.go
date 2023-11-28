package store

import (
	"errors"

	"github.com/lithammer/shortuuid/v4"
)

func NewRequest(username string) error {

	var userCount int
	err := Db_Conn.Get(&userCount, "SELECT COUNT(*) FROM User WHERE username = ?", username)
	if err != nil {
		return err
	}

	if userCount == 0 {
		return errors.New("User not found")
	}

	query1 := `
        SELECT 
            COUNT(*)
        FROM 
            Request r
        WHERE 
            r.status = 'active'
        AND 
            r.requestorId = (SELECT id FROM User WHERE username = ?)
    `

	var count int
	err = Db_Conn.Get(&count, query1, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("User have an active request")
	}

	query := `
        INSERT INTO 
            Request (id, requestorId) 
        VALUES (
            ?,
            (SELECT id FROM User WHERE username = ?)
        )
    `

	uuid := shortuuid.New()
	_, err = Db_Conn.Exec(query, uuid, username)
	if err != nil {
		return err
	}

	return nil
}
