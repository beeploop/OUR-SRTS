package store

import (
	"fmt"
	"time"
)

func UpdateFile(table, controlNumber, location string) error {
	query := fmt.Sprintf("UPDATE %v SET location = ?, updatedAt = ? WHERE owner = ?", table)

	_, err := Db_Conn.Exec(query, location, time.Now(), controlNumber)
	if err != nil {
		return err
	}

	return nil
}
