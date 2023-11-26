package types

type Submit struct {
	Search  string `form:"search" binding:"required"`
	Program string `form:"program" binding:"required"`
}
