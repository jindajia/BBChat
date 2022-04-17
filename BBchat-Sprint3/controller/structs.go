package handlers

import (
	"github.com/gorilla/websocket"
	"time"
)

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

type UserFriendsList struct {
	Username    string `json:"username"`
	Friendslist string `json:"friendslist"`
	UserID      string `json:"userId"`
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

type CreateRoomDetailResponsePayloadStruct struct {
	Username             string `json:"username"`
	UserID               string `json:"userID"`
	RoomNo               string `json:"roomNo"`
	RoomID               string `json:"roomID"`
	GenerateRoomPassword string `json:"generateRoomPassword"`
	RoomPassword         string `json:"roomPassword"`
}

type JoinRoomDetailResponsePayloadStruct struct {
	Username     string `json:"username"`
	RoomNo       string `json:"roomNo"`
	RoomPassword string `json:"roomPassword"`
}
type RoomInforStruct struct {
	RoomNo       string `json:"roomNo"`
	RoomPassword string `json:"roomPassword"`
}
type RoomDBstruct struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	RoomNo       string    `json:"roomNo"`
	Username     string    `json:"username"`
	RoomPassword string    `json:"roomPassword"`
	CreateTime   time.Time `json:"createtime"`
	RoomMember   string    `json:"roommember"`
}
