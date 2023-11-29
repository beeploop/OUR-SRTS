package store

import "github.com/BeepLoop/registrar-digitized/types"

func GetStaff() (*[]types.User, error) {
	query := `
        SELECT
            fullname, username, role, status
        FROM
            User
        WHERE
            role = 'staff'
    `

	var users []types.User
	err := Db_Conn.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
