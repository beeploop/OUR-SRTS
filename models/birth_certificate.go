package models

import "time"

type BirthCertificate struct {
	Id        int       `db:"id"`
	Location  string    `db:"location"`
	StudentId int       `db:"studentId"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}
