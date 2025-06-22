package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/BhandariG29/golang-jwt-project/middleware"
	controller "github.com/BhandariG29/golang-jwt-project/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}