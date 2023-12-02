package types

type SearchData struct {
	SearchTerm string `form:"searchterm" binding:"required"`
	Program    string `form:"program" binding:"required"`
	Type       string `form:"type" binding:"required"`
}
