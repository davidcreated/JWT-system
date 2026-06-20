package controllers

import (
	"github.com/davidcreated/golang-jwt/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// SignUP handles user registration
func SignUP() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Signup endpoint",
		})
	}
}

// Login handles user authentication
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login endpoint",
		})
	}
}

func HashPassword()

func VerifyPassword()

func GetUser()

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
	}
}
