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
			"components/nmat.html",
			"components/indigency.html",
			"components/update-other-file.html",
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
