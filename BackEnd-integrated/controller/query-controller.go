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
		if !FriendesList(userDetailsRequestPayload) {
			return "", errors.New("Request failed to complete, we are working on it")
		}

		return registrationQueryObjectID.Hex(), nil
	}
}
func FriendesList(userDetailsRequestPayload UserDetailsRequestPayloadStruct) bool {
	collection := db.MongoDBClient.Database("BBchattest").Collection("usersfriends")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	registrationQueryResponse, registrationError := collection.InsertOne(ctx, bson.M{
		"username":    userDetailsRequestPayload.Username,
		"friendslist": GetUserByUsername(userDetailsRequestPayload.Username).ID,
		"userId":      GetUserByUsername(userDetailsRequestPayload.Username).ID,
	})
	defer cancel()
	registrationQueryObjectID, registrationQueryObjectIDError := registrationQueryResponse.InsertedID.(primitive.ObjectID)

	if onlineStatusError := UpdateUserOnlineStatusByUserID(registrationQueryObjectID.Hex(), "Y"); onlineStatusError != nil {
		return false
	}

	if registrationError != nil || !registrationQueryObjectIDError {
		return false
	}

	return true

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
func StoreNewBroadcastMessages(messagePayload MessagePayloadStruct) bool {
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("broadcasts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, registrationError := collection.InsertOne(ctx, bson.M{
		"fromUserID": messagePayload.FromUserID,
		"message":    messagePayload.Message,
	})
	defer cancel()

	if registrationError == nil {
		return false
	}
	return true
}

// StoreNewChatMessages is used for storing a new message
func StoreNewDriftBottles(driftBottlePayload DriftBottlePayloadStruct) bool {
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("driftBottles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, registrationError := collection.InsertOne(ctx, bson.M{
		"fromUserID": driftBottlePayload.FromUserID,
		"message":    driftBottlePayload.Message,
		"toUserID":   driftBottlePayload.ToUserID,
	})
	defer cancel()

	if registrationError == nil {
		return false
	}
	return true
}

func StoreNewChatImages(imagePayload ImagePayloadStruct) bool {
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	_, registrationError := collection.InsertOne(ctx, bson.M{
		"fromUserID": imagePayload.FromUserID,
		"image":    imagePayload.Image,
		"toUserID":   imagePayload.ToUserID,
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
				Image:      conversation.Image,
			})
		}
	}
	return conversations
}

func GetBlindChattingBetweenTwoUsers(toUserID string, fromUserID string) []BlindChattingStruct {
	var blindChattings []BlindChattingStruct

	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("driftBottles")
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
		return blindChattings
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var blindChatting BlindChattingStruct
		err := cursor.Decode(&blindChatting)

		if err == nil {
			blindChattings = append(blindChattings, BlindChattingStruct{
				ID:         blindChatting.ID,
				FromUserID: blindChatting.FromUserID,
				ToUserID:   blindChatting.ToUserID,
				Message:    blindChatting.Message,
				Image:      blindChatting.Image,
			})
		}
	}
	return blindChattings

}

func GetBroadcast()[]BroadcastStruct {
    var broadcasts []BroadcastStruct

    collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("broadcasts")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    cursor, queryError := collection.Find(ctx,bson.D{{}})
    defer cancel()

    if queryError != nil {
        return broadcasts
    }

    for cursor.Next(context.TODO()) {
        //Create a value into which the single document can be decoded
        var broadcast BroadcastStruct
        err := cursor.Decode(&broadcast)

        if err == nil {
            broadcasts = append(broadcasts, BroadcastStruct{
                ID:         broadcast.ID,
                FromUserID: broadcast.FromUserID,
                Message:    broadcast.Message,
                Image:      broadcast.Image,
            })
        }
    }
    return broadcasts
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
	//var roomdb RoomDBstruct
	collection := db.MongoDBClient.Database(os.Getenv("MONGODB_DATABASE")).Collection("room")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	count, err := collection.CountDocuments(ctx, bson.M{"roomNo": CreateRoomDetailResponsePayload.RoomNo})
	tim := time.Now()
	//local1, _ := time.LoadLocation("Local")
	//log.Println(tim)
	//log.Println(count)
	if err != nil {
		return "", errors.New("Request failed to complete, we are working on it")
	} else {
		if count > 0 {
			cursor, _ := collection.Find(ctx, bson.M{
				"roomNo": CreateRoomDetailResponsePayload.RoomNo,
			})
			for cursor.Next(context.TODO()) {
				//Create a value into which the single document can be decoded
				var roomdbtem RoomDBstruct

				err := cursor.Decode(&roomdbtem)
				//tim1, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
				//log.Println(tim1.Sub(tim).Seconds())

				if err == nil {
					log.Println(tim.Sub(roomdbtem.CreateTime).Minutes())
					if tim.Sub(roomdbtem.CreateTime).Minutes() < 10 {

						return "", errors.New("RoomNo is not available, please pick another one")
					}
				}
			}
		} else {
			return "Yes", nil
		}

		// Sort by `price` field descending

		//_ = collection.FindOne(ctx, bson.M{
		//"roomNo": CreateRoomDetailResponsePayload.RoomNo,}).Decode(&roomdb)
		defer cancel()

	}
	return "Yes", nil

}

func CreateRoomQueryHandler(CreateRoomDetailResponsePayload CreateRoomDetailResponsePayloadStruct) (error, RoomInforStruct) {
	var RoomInfor RoomInforStruct
	var pass string
	var roomdb RoomDBstruct
	ti := time.Now()
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
	count, _ := collection.CountDocuments(ctx, bson.M{"roomNo": CreateRoomDetailResponsePayload.RoomNo})
	if count != 0 {

	}
	createRoomQueryResponse, createRoomError := collection.InsertOne(ctx, bson.M{
		"username":     CreateRoomDetailResponsePayload.Username,
		"roomPassword": newPasswordHash,
		"roomNo":       CreateRoomDetailResponsePayload.RoomNo,
		"createtime":   ti,
		"roommember":   CreateRoomDetailResponsePayload.UserID,
	})

	createroomQueryObjectID, createroomQueryObjectIDError := createRoomQueryResponse.InsertedID.(primitive.ObjectID)

	if createRoomError != nil || !createroomQueryObjectIDError {
		return errors.New("Request failed to complete, we are working on it"), RoomInfor
	} else {
		_ = collection.FindOne(ctx, bson.M{
			"_id": createroomQueryObjectID,
		}).Decode(&roomdb)
		defer cancel()
		log.Println(roomdb.CreateTime)
		return nil, RoomInforStruct{
			RoomNo:       roomdb.RoomNo,
			RoomPassword: pass,
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
				var roomDB RoomDBstruct
				err := cursor.Decode(&roomDB)
				log.Println(tim.Sub(roomDB.CreateTime).Minutes())
				if err == nil {
					if tim.Sub(roomDB.CreateTime).Minutes() < 10 {
						passwordisOkay := VerifyPassword(JoinRoomDetailResponsePayload.RoomPassword, roomDB.RoomPassword)
						if passwordisOkay == nil {
							s := strings.Split(roomDB.RoomMember, " ")
							for _, val := range s {
								if val == userId {
									return "", errors.New("user has been the member of the group chat")
								}
							}
							log.Println(roomDB.ID)
							updateUserId := roomDB.RoomMember + " " + userId
							log.Println(updateUserId)
							collection.UpdateOne(ctx, bson.M{"createtime": roomDB.CreateTime}, bson.M{"$set": bson.M{"roommember": updateUserId}})

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
		_ = collection.FindOne(ctx, bson.M{
			"roomNo": JoinRoomDetailResponsePayload.RoomNo,
		}).Decode(&roomDB)
		//log.Println(userId)
		//log.Println(tim.Sub(roomDB.CreateTime).Minutes())
		if tim.Sub(roomDB.CreateTime).Minutes() < 10 {
			passwordisOkay := VerifyPassword(JoinRoomDetailResponsePayload.RoomPassword, roomDB.RoomPassword)
			if passwordisOkay == nil {
				s := strings.Split(roomDB.RoomMember, " ")
				for _, val := range s {
					if val == userId {
						return "", errors.New("user has been the member of the group chat")
					}
				}
				updateUserId := roomDB.RoomMember + " " + userId
				log.Println(updateUserId)
				collection.UpdateOne(ctx, bson.M{"roomNo": roomDB.RoomNo}, bson.M{"$set": bson.M{"roommember": updateUserId}})
				return "user joint the group chat", nil
			} else {
				return "", errors.New("Password isn't correct")
			}
		} else {
			return "", errors.New(JoinRoomDetailResponsePayload.RoomNo + "is not available")
		}
	}
	defer cancel()
	return "", errors.New("Request failed to complete, we are working on it")

}
func CheckUser(userId string) bool {
	collection := db.MongoDBClient.Database("BBchattest").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	count, err := collection.CountDocuments(ctx, bson.M{"id": userId})
	defer cancel()
	log.Println(count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func StoreNewAddFriendsMessages(messagePayload MessagePayloadStruct) (string, error) {
	collection := db.MongoDBClient.Database("BBchattest").Collection("usersfriends")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var friendslist UserFriendsList
	cursor, queryError := collection.Find(ctx, bson.M{"userId": messagePayload.FromUserID})
	defer cancel()
	if queryError != nil {
		return "", errors.New("Request failed to complete, we are working on it")
	}
	err := cursor.Decode(&friendslist)
	if err != nil {
		return "", errors.New("Request failed to complete, we are working on it")
	}

	s := strings.Split(friendslist.Friendslist, " ")
	for _, val := range s {
		if val == messagePayload.ToUserID {
			return "Users has been your friends", errors.New(messagePayload.ToUserID + " has been your friends")
		}
	}

	return "Yes", nil

}

func UpdateFriendsList(messagePayload MessagePayloadStruct) error {
	collection := db.MongoDBClient.Database("BBchattest").Collection("usersfriends")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var friendslist UserFriendsList
	var friends UserFriendsList
	_ = collection.FindOne(ctx, bson.M{"userId": messagePayload.FromUserID}).Decode(&friendslist)
	_ = collection.FindOne(ctx, bson.M{"userId": messagePayload.ToUserID}).Decode(&friends)

	updateUserlist := friendslist.Friendslist + " " + messagePayload.ToUserID
	collection.UpdateOne(ctx, bson.M{"userId": messagePayload.FromUserID}, bson.M{"$set": bson.M{"friendslist": updateUserlist}})

	updateUserlist1 := friends.Friendslist + " " + messagePayload.FromUserID
	log.Println(updateUserlist1)
	collection.UpdateOne(ctx, bson.M{"userId": messagePayload.ToUserID}, bson.M{"$set": bson.M{"friendslist": updateUserlist1}})
	defer cancel()
	return nil

}
