package server

import (
	"fmt"

	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/types"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleManageStaff(c *gin.Context) {
	html := utils.HtmlParser(
		"admin/manage-staff.html",
		"components/head.html",
		"components/header.html",
		"components/sidebar.html",
		"components/add-staff.html",
		"components/enable-staff-modal.html",
		"components/disable-staff-modal.html",
		"components/delete-staff-modal.html",
	)

	user := utils.GetUserInSession(c)
	requestCount := store.CountActiveRequests()

	users, err := store.GetStaff()
	if err != nil {
		fmt.Println("err getting users: ", err)

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
