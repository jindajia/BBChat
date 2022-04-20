package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//func TestRenderHome(t *testing.T) {
//
//}

func TestIsUsernameAvailable(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/isUsernameAvailable/{username}", IsUsernameAvailable)

	reader := strings.NewReader(`{"Username":"legalName"}`)
	r, _ := http.NewRequest(http.MethodPost, "/isUsernameAvailable/legalName", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
	}
	nameResponse := new(APIResponseStruct)
	json.Unmarshal(w.Body.Bytes(), nameResponse)
	if (strings.Compare(nameResponse.Message, "Username is not available.")) != 0 {
		t.Errorf("The test name is not available ")
	}
}

func TestLogin(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", Login)
	reader := strings.NewReader(`{"Username": "zascauchy","Password": "kkkna784984"}`)
	r, _ := http.NewRequest(http.MethodPost, "/login", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	loginResponse := new(APIResponseStruct)
	json.Unmarshal(w.Body.Bytes(), loginResponse)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
		if strings.Compare(loginResponse.Message, "Your Login Password is incorrect.") != 0 {
			print("pw is not correct")
		}
		if strings.Compare(loginResponse.Message, "This account does not exist in our system.") != 0 {
			print("account not registered")
		}
	}

}

func TestRegistertation(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/registration", Registertation)
	reader := strings.NewReader(`{"Username": "zascauchy","Password": "kkkna784984"}`)
	r, _ := http.NewRequest(http.MethodPost, "/registration", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	regisResponse := new(APIResponseStruct)
	json.Unmarshal(w.Body.Bytes(), regisResponse)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
		if strings.Compare(regisResponse.Message, "Password can't be empty.") != 0 {
			print("Password can't be empty.")
		}
	} else {
		print("registered Successfully.")
	}
}

func TestUserSessionCheck(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/userSessionCheck/{userID}", UserSessionCheck)
	reader := strings.NewReader(`{"userID": "622261e5f61f9a3ee8b25101"}`)
	r, _ := http.NewRequest(http.MethodPost, "/userSessionCheck/{userID}", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)

	resp := w.Result()
	uscResponse := new(APIResponseStruct)
	json.Unmarshal(w.Body.Bytes(), uscResponse)

	if resp.StatusCode == http.StatusOK {
		if uscResponse.Response == false {
			print("session doesn't exist.")
		} else {
			print("session checked.")
		}
	} else {
		t.Errorf("Response code is %v", resp.StatusCode)
	}
}

//func TestGetMessagesHandler(t *testing.T) {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/registration", GetMessagesHandler)
//	reader := strings.NewReader(`{"Username": "zascauchy","Password": "kkkna784984"}`)
//	r, _ := http.NewRequest(http.MethodGet, "/registration", reader)
//	w := httptest.NewRecorder()
//	mux.ServeHTTP(w, r)
//
//	resp := w.Result()
//	regisResponse := new(APIResponseStruct)
//	json.Unmarshal(w.Body.Bytes(), regisResponse)
//
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("Response code is %v", resp.StatusCode)
//		if strings.Compare(regisResponse.Message, "Password can't be empty.") != 0 {
//			print("Password can't be empty.")
//		}
//	} else {
//		print("registered Successfully.")
//	}
//}
