package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/registrar/utils"
)

func HandleHome(c *gin.Context) {
	c.Request.Method = "GET"
	c.Redirect(http.StatusSeeOther, "/auth/login")

	// html := utils.HtmlParser("index.html", "components/head.html", "components/header.html")
	// Router.SetHTMLTemplate(html)
	//
	// c.HTML(http.StatusOK, "index.html", nil)
}
