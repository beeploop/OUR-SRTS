package server

import (
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleGetLogin(c *gin.Context) {
    logrus.Info("Hit Get login route")
	html := utils.HtmlParser("login.html", "components/head.html", "components/header.html")

	html.Execute(c.Writer, nil)
}
