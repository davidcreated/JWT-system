package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
err = nil 
	if userType != role {
		err = errors.New("unauthorized access")
	}

	return err
}

func MatchUserTypeToUid(c *gin.Context, userType string, uid string) (err error) {
	currentUserType := c.GetString("user_type")
	currentUID := c.GetString("user_id")

	if currentUserType == "USER" && currentUID != uid {
		err = errors.New("unauthorized access")
		return err

	}

	err =  CheckUserType(c, userType)
	return err
}
