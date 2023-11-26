package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/registrar/store"
	"github.com/registrar/types"
	"github.com/registrar/utils"
)

func HandleStudentRoutes(student *gin.RouterGroup) {
	student.GET("/:id", func(c *gin.Context) {
		user := utils.GetUserInSession(c)
		html := utils.HtmlParser(
			"data.html",
			"components/header.html",
			"components/editStudent.html",
			"components/update-file.html",
			"components/photo.html",
			"components/birth-certificate.html",
			"components/clearance.html",
			"components/NOAP.html",
			"components/Usepat.html",
			"components/TOR.html",
			"components/GoodMoral.html",
			"components/Form138.html",
			"components/PDS.html",
			"components/DPP.html",
			"components/HD.html",
			"components/Marriage-Certificate.html",
			"components/Promissory-Note.html",
			"components/HSD.html",
			"components/MedCert.html",
			"components/Form137.html",
			"components/ApprovalSheet.html",
			"components/AFG.html",
			"components/COLI.html",
			"components/undertaking.html",
			"components/LOA.html",
			"components/AdvancedCredit.html",
			"components/INC.html",
			"components/SubjectValidation.html",
			"components/Substitution.html",
		)

		id := c.Params.ByName("id")

		files, err := store.GetStudentFiles(id)
		if err != nil {
			fmt.Println("err getting files: ", err)
			html.Execute(c.Writer, gin.H{
				"user":  user,
				"files": files,
			})
			return
		}

		html.Execute(c.Writer, gin.H{
			"user":  user,
			"files": files,
		})
	})

	student.POST("/search", func(c *gin.Context) {

		referer := c.Request.Header.Get("Referer")
		redirectUrl := strings.Split(referer, "?")[0]

		var input types.Submit
		err := c.ShouldBindWith(&input, binding.Form)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
			return
		}

		fmt.Println("\033[35m input: ", input)

		if input.Search == "" {
			c.Request.Method = "GET"
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
			return
		}

		students, err := store.SearchStudent(input.Search, input.Program)
		if err != nil {
			c.Request.Method = "GET"
			c.Redirect(http.StatusMovedPermanently, redirectUrl+"?term="+input.Search)
			return
		}

		fmt.Println("\033[35m result: ", students)

		session := sessions.Default(c)
		session.Set("search-result", students)
		session.Save()

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, redirectUrl+"?term="+input.Search)

	})
}
