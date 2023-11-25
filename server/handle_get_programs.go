package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/registrar/store"
)

func HandleGetPrograms(c *gin.Context) {
	list, err := store.GetProgramsAndMajors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
