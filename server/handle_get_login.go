package server

import (
	"github.com/gin-gonic/gin"
	"github.com.BeepLoop/registrar-digitized/utils"
)

func HandleGetLogin(c *gin.Context) {
	html := utils.HtmlParser("login.html", "components/head.html", "components/header.html")

	html.Execute(c.Writer, nil)
}
