package server

import "github.com/gin-gonic/gin"

func HandleRequestRoutes(request *gin.RouterGroup) {

	request.GET("/", HandleRequests)

	request.POST("/reject", HandleRequestReject)

	request.POST("/fulfill", HandleRequestFulfill)

}
