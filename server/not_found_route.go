package server

import (
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleNotFoundRote(c *gin.Context) {
	html := utils.HtmlParser(
		"404.html",
	)

	html.Execute(c.Writer, nil)
}
