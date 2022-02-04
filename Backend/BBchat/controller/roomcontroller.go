package controller

import (
	"BBchat/database"
	"BBchat/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strings"
	"time"
)

var roomcollection *mongo.Collection = database.OpenCollection(database.Client, "room")
var validate1 = validator.New()

func SignRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var room models.Room

		if err := c.BindJSON(&room); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate1.Struct(room)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := roomcollection.CountDocuments(ctx, bson.M{"roomno": room.RoomNO})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			var result struct {
				//User_name string
				Email string
			}
			roomcollection.FindOne(ctx, bson.M{"roomno": room.RoomNO}).Decode(&result)
			s := strings.Split(result.Email, " ")
			for _, val := range s {
				if val == room.Email {
					c.JSON(http.StatusOK, "user has been the member of the group chat")
					return
				}
			}
			updateemail := result.Email + " " + room.Email
			roomcollection.UpdateOne(ctx, bson.M{"roomno": room.RoomNO}, bson.M{"$set": bson.M{"email": updateemail}})
			c.JSON(http.StatusOK, "user joint the group chat")
			return
		}
		fmt.Println("stop")

		room.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		room.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		room.ID = primitive.NewObjectID()
		room.Room_id = room.ID.Hex()

		InsertionNumber, insertErr := roomcollection.InsertOne(ctx, room)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, InsertionNumber)

	}
}
