package models

type User struct {
	Id       int    `db:"id"`
	Fullname string `db:"fullname"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Status   string `db:"status"`
}
