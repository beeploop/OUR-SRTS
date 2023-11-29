package store

import (
	"github.com/BeepLoop/registrar-digitized/models"
)

func GetUserInfo(username string) (models.User, error) {

	var user models.User
	err := Db_Conn.Get(&user, "SELECT * FROM User WHERE username = ?", username)
	if err != nil {
		return user, err
	}

	return user, nil
}
