package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")

	if userType != role {
		return errors.New("unauthorized access")
	}

	return nil
}

func MatchUserTypeToUid(c *gin.Context, userType string, uid string) error {
	currentUserType := c.GetString("user_type")
	currentUID := c.GetString("user_id")

	if currentUserType == "USER" && currentUID != uid {
		return errors.New("unauthorized access")
	}

	return CheckUserType(c, userType)
}
