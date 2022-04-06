package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

// RenderHome Rendering the Home Page
func RenderHome(responseWriter http.ResponseWriter, request *http.Request) {
	response := APIResponseStruct{
		Code:     http.StatusOK,
		Status:   http.StatusText(http.StatusOK),
		Message:  "This is an API for Realtime Private chat application build in GoLang",
		Response: nil,
	}
	ReturnResponse(responseWriter, request, response)
}

//IsUsernameAvailable function will handle the availability of username
func IsUsernameAvailable(responseWriter http.ResponseWriter, request *http.Request) {
	type usernameAvailableResposeStruct struct {
		IsUsernameAvailable bool `json:"isUsernameAvailable"`
	}
	var response APIResponseStruct
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	username := mux.Vars(request)["username"]

	// Checking if username is not empty & has only AlphaNumeric charecters
	if !IsAlphaNumeric(username) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		isUsernameAvailable := IsUsernameAvailableQueryHandler(username)
		if isUsernameAvailable {
			response = APIResponseStruct{
				Code:    http.StatusOK,
				Status:  http.StatusText(http.StatusOK),
				Message: "Username is available.",
				Response: usernameAvailableResposeStruct{
					IsUsernameAvailable: isUsernameAvailable,
				},
			}
		} else {
			response = APIResponseStruct{
				Code:    http.StatusOK,
				Status:  http.StatusText(http.StatusOK),
				Message: "Username is not available.",
				Response: usernameAvailableResposeStruct{
					IsUsernameAvailable: isUsernameAvailable,
				},
			}
		}
		ReturnResponse(responseWriter, request, response)
	}
}
// To be decided: whether or not we should check duplicate room names?


//Login function will login the users
func Login(responseWriter http.ResponseWriter, request *http.Request) {
	var userDetails UserDetailsRequestPayloadStruct

	decoder := json.NewDecoder(request.Body)
	requestDecoderError := decoder.Decode(&userDetails)
	defer request.Body.Close()

	if requestDecoderError != nil {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username and Password can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		if userDetails.Username == "" {
			response := APIResponseStruct{
				Code:     http.StatusBadRequest,
				Status:   http.StatusText(http.StatusBadRequest),
				Message:  "Username can't be empty.",
				Response: nil,
			}
			ReturnResponse(responseWriter, request, response)
		} else if userDetails.Password == "" {
			response := APIResponseStruct{
				Code:     http.StatusInternalServerError,
				Status:   http.StatusText(http.StatusInternalServerError),
				Message:  "Password can't be empty.",
				Response: nil,
			}
			ReturnResponse(responseWriter, request, response)
		} else {

			userDetails, loginErrorMessage := LoginQueryHandler(userDetails)

			if loginErrorMessage != nil {
				response := APIResponseStruct{
					Code:     http.StatusNotFound,
					Status:   http.StatusText(http.StatusNotFound),
					Message:  loginErrorMessage.Error(),
					Response: nil,
				}
				ReturnResponse(responseWriter, request, response)
			} else {
				response := APIResponseStruct{
					Code:     http.StatusOK,
					Status:   http.StatusText(http.StatusOK),
					Message:  "User Registration Completed.",
					Response: userDetails,
				}
				ReturnResponse(responseWriter, request, response)
			}
		}
	}
}

//Registertation function will login the users
func Registertation(responseWriter http.ResponseWriter, request *http.Request) {
	var userDetailsRequestPayload UserDetailsRequestPayloadStruct

	decoder := json.NewDecoder(request.Body)
	requestDecoderError := decoder.Decode(&userDetailsRequestPayload)
	defer request.Body.Close()

	if requestDecoderError != nil {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Request failed to complete, we are working on it",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		if userDetailsRequestPayload.Username == "" {
			response := APIResponseStruct{
				Code:     http.StatusBadRequest,
				Status:   http.StatusText(http.StatusBadRequest),
				Message:  "Username can't be empty.",
				Response: nil,
			}
			ReturnResponse(responseWriter, request, response)
		} else if userDetailsRequestPayload.Password == "" {
			response := APIResponseStruct{
				Code:     http.StatusInternalServerError,
				Status:   http.StatusText(http.StatusInternalServerError),
				Message:  "Password can't be empty.",
				Response: nil,
			}
			ReturnResponse(responseWriter, request, response)
		} else {
			userObjectID, registrationError := RegisterQueryHandler(userDetailsRequestPayload)
			if registrationError != nil {
				response := APIResponseStruct{
					Code:     http.StatusInternalServerError,
					Status:   http.StatusText(http.StatusInternalServerError),
					Message:  "Request failed to complete, we are working on it",
					Response: nil,
				}
				ReturnResponse(responseWriter, request, response)
			} else {
				response := APIResponseStruct{
					Code:    http.StatusOK,
					Status:  http.StatusText(http.StatusOK),
					Message: "User Registration Completed.",
					Response: UserDetailsResponsePayloadStruct{
						Username: userDetailsRequestPayload.Username,
						UserID:   userObjectID,
					},
				}
				ReturnResponse(responseWriter, request, response)
			}
		}
	}
}

//UserSessionCheck function will check login status of the user
func UserSessionCheck(responseWriter http.ResponseWriter, request *http.Request) {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	userID := mux.Vars(request)["userID"]

	if !IsAlphaNumeric(userID) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		uerDetails := GetUserByUserID(userID)
		if uerDetails == (UserDetailsStruct{}) {
			response := APIResponseStruct{
				Code:     http.StatusOK,
				Status:   http.StatusText(http.StatusOK),
				Message:  "You are not logged in.",
				Response: false,
			}
			ReturnResponse(responseWriter, request, response)
		} else {
			response := APIResponseStruct{
				Code:     http.StatusOK,
				Status:   http.StatusText(http.StatusOK),
				Message:  "You are logged in.",
				Response: uerDetails.Online == "Y",
			}
			ReturnResponse(responseWriter, request, response)
		}
	}
}

//GetMessagesHandler function will fetch the messages between two users
func GetDriftBottlesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	toUserID := mux.Vars(request)["toUserID"]
	fromUserID := mux.Vars(request)["fromUserID"]

	if !IsAlphaNumeric(fromUserID) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		bottles := GetBlindChattingBetweenTwoUsers(toUserID, fromUserID)
		response := APIResponseStruct{
			Code:     http.StatusOK,
			Status:   http.StatusText(http.StatusOK),
			Message:  "Username is available.",
			Response: bottles,
		}
		ReturnResponse(responseWriter, request, response)
	}
}



//GetMessagesHandler function will fetch the messages between two users
func GetMessagesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	toUserID := mux.Vars(request)["toUserID"]
	fromUserID := mux.Vars(request)["fromUserID"]

	if !IsAlphaNumeric(toUserID) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else if !IsAlphaNumeric(fromUserID) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		conversations := GetConversationBetweenTwoUsers(toUserID, fromUserID)
		response := APIResponseStruct{
			Code:     http.StatusOK,
			Status:   http.StatusText(http.StatusOK),
			Message:  "Username is available.",
			Response: conversations,
		}
		ReturnResponse(responseWriter, request, response)
	}
}

//GetMessagesHandler function will fetch the messages between two users
func GetBroadcastHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString

	fromUserID := mux.Vars(request)["fromUserID"]

	if !IsAlphaNumeric(fromUserID) {
		response := APIResponseStruct{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  "Username can't be empty.",
			Response: nil,
		}
		ReturnResponse(responseWriter, request, response)
	} else {
		broadcasts := GetBroadcast()
		response := APIResponseStruct{
			Code:     http.StatusOK,
			Status:   http.StatusText(http.StatusOK),
			Message:  "Username is available.",
			Response: broadcasts,
		}
		ReturnResponse(responseWriter, request, response)
	}
}

