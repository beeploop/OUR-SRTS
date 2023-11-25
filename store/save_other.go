package store

import "fmt"

func SaveOtherFile(filename, location, ownerNumber, table string) error {
	query := fmt.Sprintf("INSERT INTO %s (filename, location, owner) VALUES (?, ?, ?)", table)

	stmt, err := Db_Conn.Preparex(query)
	if err != nil {
		fmt.Println("err preparing sql: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(filename, location, ownerNumber)
	if err != nil {
		fmt.Println("err executing sql: ", err)
		return err
	}
	fmt.Println("executed sql")

	return nil
}
