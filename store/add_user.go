package store

import "github.com.BeepLoop/registrar-digitized/types"

func AddUser(user types.StaffInfo) error {

	query := `
        INSERT INTO
            User (fullname, username, password)
        VALUES (?, ?, ?)
    `

	_, err := Db_Conn.Exec(query, user.Fullname, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
