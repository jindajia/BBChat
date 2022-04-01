package handlers

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	db "private-chat/database"
)

// UpdateUserOnlineStatusByUserID will update the online status of the user
func UpdateUserOnlineStatusByUserID(userID string, status string) error {
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil
	}

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, queryError := collection.UpdateOne(ctx, bson.M{"_id": docID}, bson.M{"$set": bson.M{"online": status}})
	defer cancel()

	if queryError != nil {
		return errors.New("Request failed to complete, we are working on it")
	}
	return nil
}

// GetUserByUsername function will return user datails based username
func GetUserByUsername(username string) UserDetailsStruct {
	var userDetails UserDetailsStruct

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = collection.FindOne(ctx, bson.M{
		"username": username,
	}).Decode(&userDetails)

	defer cancel()

	return userDetails
}

// GetUserByUserID function will return user datails based username
func GetUserByUserID(userID string) UserDetailsStruct {
	var userDetails UserDetailsStruct

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return UserDetailsStruct{}
	}

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_ = collection.FindOne(ctx, bson.M{
		"_id": docID,
	}).Decode(&userDetails)

	defer cancel()

	return userDetails
}

// IsUsernameAvailableQueryHandler function will check username from the database
func IsUsernameAvailableQueryHandler(username string) bool {
	userDetails := GetUserByUsername(username)
	if userDetails == (UserDetailsStruct{}) {
		return true
	}
	return false
}

// LoginQueryHandler function will check username from the database
func LoginQueryHandler(userDetailsRequestPayload UserDetailsRequestPayloadStruct) (UserDetailsResponsePayloadStruct, error) {
	if userDetailsRequestPayload.Username == "" {
		return UserDetailsResponsePayloadStruct{}, errors.New("Username can't be empty.")
	} else if userDetailsRequestPayload.Password == "" {
		return UserDetailsResponsePayloadStruct{}, errors.New("Password can't be empty.")
	} else {
		userDetails := GetUserByUsername(userDetailsRequestPayload.Username)
		if userDetails == (UserDetailsStruct{}) {
			return UserDetailsResponsePayloadStruct{}, errors.New("This account does not exist in our system.")
		}

		if isPasswordOkay := VerifyPassword(userDetailsRequestPayload.Password, userDetails.Password); isPasswordOkay != nil {
			return UserDetailsResponsePayloadStruct{}, errors.New("Your Login Password is incorrect.")
		}

		if onlineStatusError := UpdateUserOnlineStatusByUserID(userDetails.ID, "Y"); onlineStatusError != nil {
			return UserDetailsResponsePayloadStruct{}, errors.New("Your Login Password is incorrect.")
		}

		return UserDetailsResponsePayloadStruct{
			UserID:   userDetails.ID,
			Username: userDetails.Username,
		}, nil
	}
}

// RegisterQueryHandler function will check username from the database
func RegisterQueryHandler(userDetailsRequestPayload UserDetailsRequestPayloadStruct) (string, error) {
	if userDetailsRequestPayload.Username == "" {
		return "", errors.New("Username can't be empty.")
	} else if userDetailsRequestPayload.Password == "" {
		return "", errors.New("Password can't be empty.")
	} else {
		newPasswordHash, newPasswordHashError := HashPassword(userDetailsRequestPayload.Password)
		if newPasswordHashError != nil {
			return "", errors.New("Request failed to complete, we are working on it")
		}
		collection := db.MongoDBClient.Database("BBchattest").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		registrationQueryResponse, registrationError := collection.InsertOne(ctx, bson.M{
			"username": userDetailsRequestPayload.Username,
			"password": newPasswordHash,
			"online":   "N",
		})
		defer cancel()

		registrationQueryObjectID, registrationQueryObjectIDError := registrationQueryResponse.InsertedID.(primitive.ObjectID)

		if onlineStatusError := UpdateUserOnlineStatusByUserID(registrationQueryObjectID.Hex(), "Y"); onlineStatusError != nil {
			return " ", errors.New("Request failed to complete, we are working on it")
		}

		if registrationError != nil || !registrationQueryObjectIDError {
			return "", errors.New("Request failed to complete, we are working on it")
		}

		return registrationQueryObjectID.Hex(), nil
	}
}

// GetAllOnlineUsers function will return the all online users
func GetAllOnlineUsers(userID string) []UserDetailsResponsePayloadStruct {
	var onlineUsers []UserDetailsResponsePayloadStruct

	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return onlineUsers
	}

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, queryError := collection.Find(ctx, bson.M{
		"online": "Y",
		"_id": bson.M{
			"$ne": docID,
		},
	})
	defer cancel()

	if queryError != nil {
		return onlineUsers
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var singleOnlineUser UserDetailsStruct
		err := cursor.Decode(&singleOnlineUser)

		if err == nil {
			onlineUsers = append(onlineUsers, UserDetailsResponsePayloadStruct{
				UserID:   singleOnlineUser.ID,
				Online:   singleOnlineUser.Online,
				Username: singleOnlineUser.Username,
			})
		}
	}

	return onlineUsers
}

// StoreNewChatMessages is used for storing a new message
func StoreNewChatMessages(messagePayload MessagePayloadStruct) bool {
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, registrationError := collection.InsertOne(ctx, bson.M{
		"fromUserID": messagePayload.FromUserID,
		"message":    messagePayload.Message,
		"toUserID":   messagePayload.ToUserID,
	})
	defer cancel()

	if registrationError == nil {
		return false
	}
	return true
}

// GetConversationBetweenTwoUsers will be used to fetch the conversation between two users
func GetConversationBetweenTwoUsers(toUserID string, fromUserID string) []ConversationStruct {
	var conversations []ConversationStruct

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	queryCondition := bson.M{
		"$or": []bson.M{
			{
				"$and": []bson.M{
					{
						"toUserID": toUserID,
					},
					{
						"fromUserID": fromUserID,
					},
				},
			},
			{
				"$and": []bson.M{
					{
						"toUserID": fromUserID,
					},
					{
						"fromUserID": toUserID,
					},
				},
			},
		},
	}

	cursor, queryError := collection.Find(ctx, queryCondition)
	defer cancel()

	if queryError != nil {
		return conversations
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var conversation ConversationStruct
		err := cursor.Decode(&conversation)

		if err == nil {
			conversations = append(conversations, ConversationStruct{
				ID:         conversation.ID,
				FromUserID: conversation.FromUserID,
				ToUserID:   conversation.ToUserID,
				Message:    conversation.Message,
			})
		}
	}
	return conversations
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
		return "", errors.New("Error occurred while creating a Hash")
	}

	return string(bytes), nil
}

// ComparePasswords will create password using bcrypt
func VerifyPassword(password string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(password))
	if err != nil {
		fmt.Sprintf("login or passowrd is incorrect")
		return errors.New("The '" + password + "' and '" + providedPassword + "' strings don't match")
	}
	return nil
}
func RoomNoavailableQueryHandle(CreateRoomDetailResponsePayload CreateRoomDetailResponsePayloadStruct) (string, error) {
	var roomdb RoomDBstruct
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("room")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	count, err := collection.CountDocuments(ctx, bson.M{"roomNo": CreateRoomDetailResponsePayload.RoomNo})
	if err != nil {
		return "", errors.New("Request failed to complete, we are working on it")
	} else {
		if count > 0 {
			_ = collection.FindOne(ctx, bson.M{
				"username": CreateRoomDetailResponsePayload.RoomNo,
			}).Decode(&roomdb)
			defer cancel()

			tim := time.Now()
			if tim.Sub(roomdb.CreateTime).Minutes() < 10 {
				return "", errors.New("RoomNo is not available, please pick another one")
			} else {
				return "Yes", nil
			}
		} else {
			return "Yes", nil
		}
	}

}

func CreateRoomQueryHandler(CreateRoomDetailResponsePayload CreateRoomDetailResponsePayloadStruct) (error, RoomInforStruct) {
	var RoomInfor RoomInforStruct
	var pass string
	var roomdb RoomDBstruct
	if CreateRoomDetailResponsePayload.GenerateRoomPassword == "Yes" {
		pass = generatepassword()
	} else {
		pass = CreateRoomDetailResponsePayload.RoomPassword
	}
	newPasswordHash, newPasswordHashError := HashPassword(pass)
	if newPasswordHashError != nil {
		return errors.New("Request failed to complete, we are working on it"), RoomInfor
	}

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("room")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	createRoomQueryResponse, createRoomError := collection.InsertOne(ctx, bson.M{
		"username":     CreateRoomDetailResponsePayload.Username,
		"roomPassword": newPasswordHash,
		"roomNo":       CreateRoomDetailResponsePayload.RoomNo,
		"creat_time":   time.Parse(time.RFC3339, time.Now().Format(time.RFC3339)),
		"room_member":  CreateRoomDetailResponsePayload.UserID,
	})

	createroomQueryObjectID, createroomQueryObjectIDError := createRoomQueryResponse.InsertedID.(primitive.ObjectID)

	if createRoomError != nil || !createroomQueryObjectIDError {
		return errors.New("Request failed to complete, we are working on it"), RoomInfor
	} else {
		_ = collection.FindOne(ctx, bson.M{
			"_id": createroomQueryObjectID.Hex(),
		}).Decode(&roomdb)
		defer cancel()
		return nil, RoomInforStruct{
			RoomNo:       roomdb.RoomNo,
			RoomPassword: roomdb.RoomPassword,
			RoomID:       roomdb.RoomID,
		}
	}

}

func JoinRoomQueryHandler(JoinRoomDetailResponsePayload JoinRoomDetailResponsePayloadStruct) (string, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("room")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var roomDB RoomDBstruct
	count, err := collection.CountDocuments(ctx, bson.M{"roomNo": JoinRoomDetailResponsePayload.RoomNo})
	cursor, queryError := collection.Find(ctx, bson.M{"roomNo": JoinRoomDetailResponsePayload.RoomNo})
	tim := time.Now()
	userId := GetUserByUsername(JoinRoomDetailResponsePayload.Username).ID
	if err != nil {
		return "", errors.New("Request failed to complete, we are working on it")
	}
	if count > 1 {

		if queryError != nil {
			return "", errors.New("Request failed to complete, we are working on it")
		} else {
			for cursor.Next(context.TODO()) {

				err := cursor.Decode(&roomDB)

				if err == nil {
					if tim.Sub(roomDB.CreateTime) < 10 {
						passwordisOkay := VerifyPassword(JoinRoomDetailResponsePayload.RoomPassword, roomDB.RoomPassword)
						if passwordisOkay == nil {
							s := strings.Split(roomDB.RoomMember, " ")
							for _, val := range s {
								if val == userId {
									return "", errors.New("user has been the member of the group chat")
								}
							}
							updateUserId := roomDB.RoomMember + " " + userId
							collection.UpdateOne(ctx, bson.M{"roomID": roomDB.RoomID}, bson.M{"$set": bson.M{"room_member": updateUserId}})
							return "user joint the group chat", nil
						} else {
							return "", errors.New("Password isn't correct")
						}
					}

				} else {
					return "", errors.New("Request failed to complete, we are working on it")
				}
			}
			return "", errors.New(JoinRoomDetailResponsePayload.RoomNo + "is not available")
		}

	} else if count == 0 {
		return "", errors.New(JoinRoomDetailResponsePayload.RoomNo + "doesn't exist, please check it")
	} else {
		err := cursor.Decode(&roomDB)

		if err == nil {
			if tim.Sub(roomDB.CreateTime) < 10 {
				passwordisOkay := VerifyPassword(JoinRoomDetailResponsePayload.RoomPassword, roomDB.RoomPassword)
				if passwordisOkay == nil {
					s := strings.Split(roomDB.RoomMember, " ")
					for _, val := range s {
						if val == userId {
							return "", errors.New("user has been the member of the group chat")
						}
					}
					updateUserId := roomDB.RoomMember + " " + userId
					collection.UpdateOne(ctx, bson.M{"roomID": roomDB.RoomID}, bson.M{"$set": bson.M{"room_member": updateUserId}})
					return "user joint the group chat", nil
				} else {
					return "", errors.New("Password isn't correct")
				}
			} else {
				return "", errors.New(JoinRoomDetailResponsePayload.RoomNo + "is not available")
			}
		} else {
			return "", errors.New("Request failed to complete, we are working on it")
		}

	}
	defer cancel()
	return "", errors.New("Request failed to complete, we are working on it")

}
