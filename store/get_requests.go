package store

import (
	"fmt"

	"github.com/BeepLoop/registrar-digitized/models"
)

func GetRequests() ([]models.Request, error) {
	query := `
        SELECT 
            r.id as id,
            r.status as status,
            u.username as requestor
        FROM Request r
        LEFT JOIN User u on r.requestorId = u.id
    `

	var requests []models.Request
	err := Db_Conn.Select(&requests, query)
	if err != nil {
		fmt.Println("err getting requests: ", err)
		return nil, err
	}

	return requests, nil
}
