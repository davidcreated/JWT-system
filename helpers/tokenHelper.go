package helpers

import (
	"log"
	"os"
	"time"

	"github.com/davidcreated/golang-jwt/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email              string
	First_name         string
	Last_name          string
	Uid                string
	User_type          string
	jwt.StandardClaims // this means that we are embedding the StandardClaims struct from the jwt package into our SignedDetails struct. This allows us to include standard JWT claims (like expiration time, issuer, etc.) along with our custom claims (like email, first name, last name, etc.) in the same JWT token.
}

// this is used to open the user collection in the database. It uses the OpenCollection function from the database package to get a reference to the "user" collection in the MongoDB database.
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, uid string, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// this means
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
