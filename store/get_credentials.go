package store

type Credential struct {
	Fullname string `db:"fullname"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Status   string `db:"status"`
}

func GetCredentials(username string) (*Credential, error) {
	query := "SELECT fullname, username, password, role, status from User WHERE username=?"

	var credential Credential
	err := Db_Conn.Get(&credential, query, username)
	if err != nil {
		return nil, err
	}

	return &credential, nil
}
