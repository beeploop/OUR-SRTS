package server

import (
	"github.com/BeepLoop/registrar-digitized/config"
	"github.com/BeepLoop/registrar-digitized/store"
	"github.com/BeepLoop/registrar-digitized/utils"
	"github.com/gin-gonic/gin"
)

func HandleStudentRoutes(student *gin.RouterGroup) {
	student.GET("/:id", func(c *gin.Context) {
		user := utils.GetUserInSession(c)
		html := utils.HtmlParser(
			"data.tmpl",
			"components/header.tmpl",
			"components/editStudent.tmpl",
			"components/update-file.tmpl",
			"components/photo.tmpl",
			"components/birth-certificate.tmpl",
			"components/clearance.tmpl",
			"components/NOAP.tmpl",
			"components/Usepat.tmpl",
			"components/TOR.tmpl",
			"components/GoodMoral.tmpl",
			"components/Form138.tmpl",
			"components/PDS.tmpl",
			"components/DPP.tmpl",
			"components/HD.tmpl",
			"components/Marriage-Certificate.tmpl",
			"components/Promissory-Note.tmpl",
			"components/HSD.tmpl",
			"components/MedCert.tmpl",
			"components/Form137.tmpl",
			"components/ApprovalSheet.tmpl",
			"components/AFG.tmpl",
			"components/COLI.tmpl",
			"components/undertaking.tmpl",
			"components/LOA.tmpl",
			"components/AdvancedCredit.tmpl",
			"components/INC.tmpl",
			"components/SubjectValidation.tmpl",
			"components/Substitution.tmpl",
			"components/nmat.tmpl",
			"components/indigency.tmpl",
			"components/update-other-file.tmpl",
		)

		id := c.Params.ByName("id")

		files, err := store.GetStudentFiles(id)
		if err != nil {
			html.Execute(c.Writer, gin.H{
				"user":  user,
				"files": files,
			})
			return
		}

		localAddr := "http://" + config.Env.LocalAddr + config.Env.Port
		files.LocalAddr = localAddr

		html.Execute(c.Writer, gin.H{
			"user":  user,
			"files": files,
		})
	})
}
