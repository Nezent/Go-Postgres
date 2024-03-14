package routes

import (
	"github.com/Nezent/Go-PSQL/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUser)
	router.POST("/",controller.CreateUser)
	router.DELETE("/:id",controller.DeleteUser)
	router.PUT("/:id",controller.UpdateUser)
}