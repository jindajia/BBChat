package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/http/httptest"
	"private-chat/controller"
	"strings"
	"testing"
)

func TestRegistertation(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"passwordempty", `{"Username": "liwenzhou"}`, `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`},
		{"usernameempty", `{"Password": "liwenzhou"}`, `{"code":400,"status":"Bad Request","message":"Username can't be empty.","response":null}`},
	}
	//json := strings.NewReader(`{"Username":"1"}`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req, err = http.NewRequest("POST", "/registration", strings.NewReader(tt.param))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.Registertation)

			handler.ServeHTTP(rr, req)

			// Check the response body is what we expect.
			//expected := `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`
			if rr.Body.String() != tt.expect {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expect)
			}

		})
	}
}

func TestLogin(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"passwordempty", `{"Username": "liwenzhou"}`, `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`},
		{"usernameempty", `{"Password": "liwenzhou"}`, `{"code":400,"status":"Bad Request","message":"Username can't be empty.","response":null}`},
	}
	//json := strings.NewReader(`{"Username":"1"}`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req, err = http.NewRequest("POST", "/login", strings.NewReader(tt.param))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.Login)

			handler.ServeHTTP(rr, req)

			// Check the response body is what we expect.
			//expected := `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`
			if rr.Body.String() != tt.expect {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expect)
			}

		})
	}
}

func TestCreateroom(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tests := []struct {
		name   string
		param  string
		expect string
	}{

		{"roomnoempty", `{"Username":"kkk","RoomName":"","GenerateRoomPassword":"No","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"RoomName can't be empty.","response":null}`},
		{"generatepasswordempty", `{"Username":"kkk","RoomName":"2345","GenerateRoomPassword":"","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"You need to decide whether generate password or not.","response":null}`},
		{"Pawwordwrong", `{"Username":"kkk","RoomName":"45","GenerateRoomPassword":"No","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"Password cannot be empty.","response":null}`},
	}
	//json := strings.NewReader(`{"Username":"1"}`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req, err = http.NewRequest("POST", "/CreateRoom", strings.NewReader(tt.param))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreatRoom)

			handler.ServeHTTP(rr, req)

			// Check the response body is what we expect.
			//expected := `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`
			if rr.Body.String() != tt.expect {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expect)
			}

		})
	}
}

func TestJoineroom(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tests := []struct {
		name   string
		param  string
		expect string
	}{

		{"roomnoempty", `{"Username":"kkk","RoomNo":"","RoomPassword": "klm"}`, `{"code":400,"status":"Bad Request","message":"RoomNo can't be empty.","response":null}`},
		{"passwordempty", `{"Username":"kkk","RoomNo":"345","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"Room Password can't be empty.","response":null}`},
	}
	//json := strings.NewReader(`{"Username":"1"}`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req, err = http.NewRequest("POST", "/JoinRoom", strings.NewReader(tt.param))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.JoinRoom)

			handler.ServeHTTP(rr, req)

			// Check the response body is what we expect.
			//expected := `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`
			if rr.Body.String() != tt.expect {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expect)
			}

		})
	}
}

type SocketEventStruct struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

type EventPayLoad struct {
	FromUserID string `json:"fromUserID"`
	ToUserID   string `json:"toUserID"`
	Message    string `json:"message"`
}

func TestWebSocket(t *testing.T) {

	event := EventPayLoad{
		FromUserID: "625f5e1587dbe1be871557f2",
		ToUserID:   "625f586544c7dff685f96069",
		Message:    "this is a test message from A to B",
	}

	stu := SocketEventStruct{
		EventName:    "message",
		EventPayload: event,
	}

	marshal, _ := json.Marshal(stu)
	log.Println(string(marshal))
	url := "ws://localhost:8000/ws/625f586544c7dff685f96069"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	log.Printf("response:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Receive Message: %s", message)

	}
	<-done
}

func TestWebSocket(t *testing.T) {

	event := EventPayLoad{
		FromUserID: "625f586544c7dff685f96069",
		ToUserID:   "625f5e1587dbe1be871557f2",
		Message:    "this is a test message from B to A",
	}

	stu := SocketEventStruct{
		EventName:    "message",
		EventPayload: event,
	}

	marshal, _ := json.Marshal(stu)
	log.Println(string(marshal))
	url := "ws://localhost:8000/ws/625f5e1587dbe1be871557f2"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	log.Printf("response:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Receive Message: %s", message)

	}
	<-done
}

type BroadEventPayLoad struct {
	FromUserID string `json:"fromUserID"`
	Message    string `json:"message"`
}
func TestWebSocket(t *testing.T) {

	event := BroadEventPayLoad{
		FromUserID: "625f586544c7dff685f96069",
		Message:    "This is a broadcast test!",
	}

	stu := SocketEventStruct{
		EventName:    "broadcast",
		BroadEventPayLoad: event,
	}

	marshal, _ := json.Marshal(stu)
	log.Println(string(marshal))
	url := "ws://localhost:8000/ws/625f586544c7dff685f96069"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	log.Printf("response:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Receive Message: %s", message)

	}
	<-done
}

func TestWebSocket(t *testing.T) {

	event := EventPayLoad{
		FromUserID: "625f586544c7dff685f96069",
		ToUserID:   "625f5e1587dbe1be871557f2",
		Message:    "This is a driftBottle message",
	}

	stu := SocketEventStruct{
		EventName:    "driftBottle",
		EventPayload: event,
	}

	marshal, _ := json.Marshal(stu)
	log.Println(string(marshal))
	url := "ws://localhost:8000/ws/625f5e1587dbe1be871557f2"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	log.Printf("response:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Receive Message: %s", message)
	}
	<-done
}

