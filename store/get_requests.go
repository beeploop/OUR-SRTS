package store

import (
	"github.com/BeepLoop/registrar-digitized/models"
	"github.com/sirupsen/logrus"
)

func GetRequests() ([]models.Request, error) {
	query := `
        SELECT 
            r.id as id,
            r.status as status,
            u.username as requestor
        FROM Request r
        LEFT JOIN User u on r.requestorId = u.id
        WHERE r.status = 'active'
    `

	var requests []models.Request
	err := Db_Conn.Select(&requests, query)
	if err != nil {
		logrus.Warn("err getting requests: ", err)
		return nil, err
	}

	return requests, nil
}
