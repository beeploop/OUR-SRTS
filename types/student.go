package types

import "database/sql"

type Student struct {
	Id             int           `db:"id"`
	Control_Number string        `db:"controlNumber"`
	Lastname       string        `db:"lastname"`
	Firstname      string        `db:"firstname"`
	Middlename     string        `db:"middlename"`
	Type           string        `db:"type"`
	CivilStatus    any           `db:"civilStatus"`
	FileLocation   string        `db:"fileLocation"`
	ProgramId      sql.NullInt64 `db:"programId"`
	MajorId        sql.NullInt64 `db:"majorId"`
	CreatedAt      string        `db:"createdAt"`
	UpdatedAt      string        `db:"updatedAt"`
}

type SearchResult struct {
	Control_Number string `db:"controlNumber"`
	Lastname       string `db:"lastname"`
	Firstname      string `db:"firstname"`
	Middlename     string `db:"middlename"`
}
