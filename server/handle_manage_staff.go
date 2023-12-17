package server

import (
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleManageStaff(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/manage-staff.tmpl",
		"components/head.tmpl",
		"components/header.tmpl",
		"components/sidebar.tmpl",
		"components/add-staff.tmpl",
		"components/enable-staff-modal.tmpl",
		"components/disable-staff-modal.tmpl",
		"components/delete-staff-modal.tmpl",
	)

	user := utils.GetUserInSession(c)
	requestCount := store.CountActiveRequests()

	users, err := store.GetStaff()
	if err != nil {
		logrus.Warn("err getting staff: ", err)

		html.Execute(c.Writer, gin.H{
			"user":         user.Fullname,
			"users":        []types.User{},
			"requestCount": requestCount,
		})
		return
	}

	html.Execute(c.Writer, gin.H{
		"user":         user,
		"users":        users,
		"requestCount": requestCount,
	})
}
