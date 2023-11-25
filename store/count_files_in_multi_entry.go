package store

import "fmt"

func CountFilesInMultiEntry(table, controlNumber string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE owner = ?", table)

	var count int
	err := Db_Conn.Get(&count, query, controlNumber)
	if err != nil {
		return -1, err
	}

	return count, nil
}
