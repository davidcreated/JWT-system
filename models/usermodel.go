package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id"`
	FirstName   string             `json:"first_name" validate:"required,min=2,max=100"`
	LastName    string             `json:"last_name" validate:"required,min=2,max=100"`
	Email       string             `json:"email" validate:"required,email"`
	Password    string             `json:"password" validate:"required,min=6"`
	PhoneNumber string             `json:"phone_number" validate:"required,min=10,max=15"`
	Token       string             `json:"token"`
	UserType    string             `json:"user_type" validate:" required, eq=ADMIN|USER"`
	RefreshToken string            `json:"refresh_token"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
    UserID string `json:"user_id"`
}