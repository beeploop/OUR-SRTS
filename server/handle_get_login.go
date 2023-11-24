package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/registrar/utils"
)

func HandleGetLogin(c *gin.Context) {
	html := utils.HtmlParser("login.html", "components/head.html", "components/header.html")
	Router.SetHTMLTemplate(html)

	c.HTML(http.StatusOK, "login.html", nil)
}
