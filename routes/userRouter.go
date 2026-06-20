package routes

import (
	controller "github.com/davidcreated/golang-jwt/controllers"
	"github.com/davidcreated/golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
