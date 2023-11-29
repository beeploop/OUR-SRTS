package types

type StudentInfo struct {
	Firstname     string `form:"firstname" binding:"required"`
	Lastname      string `form:"lastname" binding:"required"`
	Middlename    string `form:"middlename" binding:"required"`
	ControlNumber string `form:"controlNumber" binding:"required"`
	FileLocaion   string `form:"fileLocation" binding:"required"`
	Type          string `form:"type" binding:"required"`
	CivilStatus   string `form:"civilStatus" binding:"required"`
	Program       string `form:"program" binding:"required"`
	Major         string `form:"major" binding:""`
}
