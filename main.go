package main

import (
	"os"
	"github.com/gin-gonic/gin"
	routes "github.com/davidcreated/golang-jwt/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())



	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Access Granted for api-1",
		})
	})



	router.GET("/api-2/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Access Granted for api-2",
		})
	})

	router.Run(":"+port)
}
