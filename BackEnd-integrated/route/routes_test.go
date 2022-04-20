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
		{"passwordempty", `{"Username": "kexinzhang"}`, `{"code":500,"status":"Internal Server Error","message":"Password can't be empty.","response":null}`},
		{"usernameempty", `{"Password": "kexinzhang"}`, `{"code":400,"status":"Bad Request","message":"Username can't be empty.","response":null}`},
	}
	//json := strings.NewReader(`{"Username":"1"}`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req, err = http.NewRequest("POST", "http://localhost:8000/registration", strings.NewReader(tt.param))
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

func TestAddFriendsWebSocket(t *testing.T) {

	event := EventPayLoad{
		FromUserID: "6260503a1190410dbe5baba8",
		ToUserID:   "626050421190410dbe5babaa",
		Message:    "I love",
	}

	stu := SocketEventStruct{
		EventName:    "add_friends",
		EventPayload: event,
	}

	event1 := EventPayLoad{
		ToUserID:   "626050421190410dbe5babaa",
		FromUserID: "6260503a1190410dbe5baba8",
		Message:    "Yes",
	}

	stu1 := SocketEventStruct{
		EventName:    "response_friends",
		EventPayload: event1,
	}
	event2 := EventPayLoad{
		FromUserID: "6260503a1190410dbe5baba8",
		ToUserID:   "626050421190410dbe5babaa",
		Message:    "I love u",
	}

	stu2 := SocketEventStruct{
		EventName:    "message",
		EventPayload: event2,
	}
	event3 := EventPayLoad{
		ToUserID:   "626050421190410dbe5babaa",
		FromUserID: "6260503a1190410dbe5baba8",
		Message:    "I love u too",
	}

	stu3 := SocketEventStruct{
		EventName:    "message",
		EventPayload: event3,
	}

	event4 := EventPayLoad{
		FromUserID: "6260503a1190410dbe5baba8",
		ToUserID:   "87071192",
		Message:    "HI, I am Kexin Zhang",
	}

	stu4 := SocketEventStruct{
		EventName:    "room-chat",
		EventPayload: event4,
	}
	event5 := EventPayLoad{
		FromUserID: "626050421190410dbe5babaa",
		ToUserID:   "87071192",
		Message:    "Hi!!",
	}

	stu5 := SocketEventStruct{
		EventName:    "room-chat",
		EventPayload: event5,
	}

	event6 := EventPayLoad{
		ToUserID:   "6260504c1190410dbe5babac",
		FromUserID: "6260503a1190410dbe5baba8",
		Message:    "Hi",
	}

	stu6 := SocketEventStruct{
		EventName:    "driftBottle",
		EventPayload: event6,
	}
	url := "ws://localhost:8000/ws/6260503a1190410dbe5baba8"
	url2 := "ws://localhost:8000/ws/626050421190410dbe5babaa"
	url3 := "ws://localhost:8000/ws/6260504c1190410dbe5babac"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	w, ress, error := websocket.DefaultDialer.Dial(url2, nil)
	r, resss, _ := websocket.DefaultDialer.Dial(url3, nil)
	if err != nil {
		log.Fatal("Connect error:", err)
	}
	if error != nil {
		log.Fatal("Connect error:", error)
	}
	log.Printf("Reponse:%s", fmt.Sprint(res))
	log.Printf("Reponse:%s", fmt.Sprint(ress))
	log.Printf("Reponse:%s", fmt.Sprint(resss))
	defer c.Close()
	defer w.Close()
	defer r.Close()
	done := make(chan struct{})
	marshal, _ := json.Marshal(stu)
	err = c.WriteMessage(websocket.TextMessage, marshal)
	marshal1, _ := json.Marshal(stu1)
	error = w.WriteMessage(websocket.TextMessage, marshal1)
	marshal2, _ := json.Marshal(stu2)
	err = c.WriteMessage(websocket.TextMessage, marshal2)
	marshal3, _ := json.Marshal(stu3)
	error = w.WriteMessage(websocket.TextMessage, marshal3)
	marshal4, _ := json.Marshal(stu4)
	err = c.WriteMessage(websocket.TextMessage, marshal4)
	marshal5, _ := json.Marshal(stu5)
	error = w.WriteMessage(websocket.TextMessage, marshal5)
	marshal6, _ := json.Marshal(stu6)
	error = r.WriteMessage(websocket.TextMessage, marshal6)
	if err != nil {
		fmt.Println(err)
	}
	if error != nil {
		fmt.Println(error)
	}
	for {
		_, message, err := w.ReadMessage()
		_, mess, error := c.ReadMessage()
		_, mes, _ := r.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		if error != nil {
			log.Fatal(err)
			break
		}
		log.Printf("Received message from user1: %s", message)
		log.Printf("Received message from user2: %s", mess)
		log.Printf("Received message from user3: %s", mes)
	}
	<-done
}
