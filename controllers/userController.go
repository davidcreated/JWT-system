package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davidcreated/golang-jwt/database"
	"github.com/davidcreated/golang-jwt/helpers"
	"github.com/davidcreated/golang-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// SignUP handles user registration
func SignUP() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationEerr := validate.Struct(user)
		if validationEerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationEerr.Error()})
			return
		}

		// count is used against validation to check if the email or phone number already exists in the database
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		}

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.PhoneNumber})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone"})
		}

		// this is the usecase of count
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error ": "this email or phone number already exists"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// package from helpers used to generate the token and refresh token for the user
		token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.User_id, *user.User_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Token = &token
		user.Refresh_token = &refreshToken

		// create user in database upon account creation
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOk, resultInsertionNumber)
	}
}

// Login handles user authentication
func Login() gin.HandlerFunc {
return func(c *gin.Context) {
var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
var user models.User 

// for existing users in db
var foundUser models.User 

c.BindJSON(&user); err != nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
defer cancel()
if err != nil{
	c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
	return
}

passwordIsValid, msg := helpers.VerifyPassword(*user.Password, *foundUser.Password)
if !passwordIsValid || msg != ""{
	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	return
}

}
}

func HashPassword()

func VerifyPassword(userPassword string , providedPassword	string)(bool, string ){
	err:= bcrypt.ComapareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	
	if err != nil {
		return false, "email or password is incorrect"
	}
	return true, ""
}

func GetUser()

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, "ADMIN", userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User

		// decode is used to convert the data from the database into a Go struct i. json response to the langauage go can read given it doesnt understand json response . hence why we created a user struct
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
