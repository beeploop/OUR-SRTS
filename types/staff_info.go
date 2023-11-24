package types

type StaffInfo struct {
	Fullname string `form:"fullname" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
