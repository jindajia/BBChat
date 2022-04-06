package route

import (
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

		{"roomnoempty", `{"Username":"kkk","RoomNo":"","GenerateRoomPassword":"No","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"RoomNo can't be empty.","response":null}`},
		{"generatepasswordempty", `{"Username":"kkk","RoomNo":"2345","GenerateRoomPassword":"","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"You need to decide whether generate password or not.","response":null}`},
		{"Pawwordwrong", `{"Username":"kkk","RoomNo":"45","GenerateRoomPassword":"No","RoomPassword": ""}`, `{"code":400,"status":"Bad Request","message":"Password cannot be empty.","response":null}`},
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
