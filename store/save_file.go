package store

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func SaveFile(location, ownerNumber, table string) error {
	query := fmt.Sprintf("INSERT INTO %s (location, owner) VALUES (?, ?)", table)

	stmt, err := Db_Conn.Preparex(query)
	if err != nil {
		logrus.Warn("err preparing sql: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(location, ownerNumber)
	if err != nil {
		logrus.Warn("err executing sql: ", err)
		return err
	}

	return nil
}
