package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type Room struct {
	ID     primitive.ObjectID `bson:"id"`
	RoomNO *string            `json:"roomno" validate:"required,min=1"`
	//Password      *string            `json:"Password" validate:"required,min=6""`
	//User_name      string   `json:"user_name" validate:"required,min=1"`
	Email string `json:"email" validate:"required,min=6"`
	//Token         *string   `json:"token"`
	User_type *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER""`
	//Refresh_token *string   `json:"refresh_token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Room_id    string    `json:"room_id"`
}
