package server

import "github.com/gin-gonic/gin"

func HandleAuthRoutes(auth *gin.RouterGroup) {

    auth.GET("/login", HandleGetLogin)

    auth.POST("/login", HandlePostLogin)

    auth.POST("/logout", HandleLogout)
}
