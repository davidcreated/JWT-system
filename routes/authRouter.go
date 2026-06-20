package routes

import (
	controller "github.com/davidcreated/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.SignUP())
	incomingRoutes.POST("users/login", controller.Login())
}
