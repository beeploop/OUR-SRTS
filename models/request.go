package models

type Request struct {
	Id        string `db:"id"`
	Requestor string `db:"requestor"`
	Status    string `db:"status"`
}
