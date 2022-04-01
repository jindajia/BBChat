package handlers

import "github.com/gorilla/websocket"

// UserDetailsStruct is a universal struct for mapping the user details
type UserDetailsStruct struct {
	ID       string `bson:"_id,omitempty"`
	Username string
	Password string
	Online   string
	SocketID string
}

// ConversationStruct is a universal struct for mapping the conversations
type ConversationStruct struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Message    string `json:"message"`
	ToUserID   string `json:"toUserID"`
	FromUserID string `json:"fromUserID"`
}

// UserDetailsRequestPayloadStruct represents payload for Login and Registration request
type UserDetailsRequestPayloadStruct struct {
	Username string
	Password string
}

// UserDetailsResponsePayloadStruct represents payload for Login and Registration response
type UserDetailsResponsePayloadStruct struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	Online   string `json:"online"`
}

// SocketEventStruct struct of socket events
type SocketEventStruct struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub                 *Hub
	webSocketConnection *websocket.Conn
	send                chan SocketEventStruct
	userID              string
}

// MessagePayloadStruct is a struct used for message Payload
type MessagePayloadStruct struct {
	FromUserID string `json:"fromUserID"`
	ToUserID   string `json:"toUserID"`
	Message    string `json:"message"`
}
// ImagePayloadStruct is a struct used for image Payload
// Image size should be smaller than 4M, it is actually similar with message.
// But we need to distinguish it from message. So that the front could know which kind of data need to be transfered into image formation
type ImagePayloadStruct struct {
	FromUserID string `json:"fromUserID"`
	ToUserID   string `json:"toUserID"`
	Image    string `json:"image"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	User_name     *string            `json:"user_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6""`
	Email         *string            `json:"email" validate:"email,required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER""`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

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