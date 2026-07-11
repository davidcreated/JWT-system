package routes

import (
	controller "github.com/davidcreated/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controller.SignUP())
	router.POST("users/login", controller.Login())
}
