package middleware

import (
	"fmt"
	"net/http"

	"github.com/davidcreated/golang-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("no authentication token provided")})
			c.Abort()
			return
		}

		claims, msg := helpers.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("token is not valid%v", msg)})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()

	}
}
