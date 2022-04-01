package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var m map[string]Hub

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	log.Println("Loadeding Routes...")

	hub := NewHub()
	go hub.Run()

	route.HandleFunc("/", RenderHome)

	route.HandleFunc("/isUsernameAvailable/{username}", IsUsernameAvailable)

	route.HandleFunc("/login", Login).Methods("POST")

	route.HandleFunc("/registration", Registertation).Methods("POST")

	route.HandleFunc("/userSessionCheck/{userID}", UserSessionCheck)

	route.HandleFunc("/getConversation/{toUserID}/{fromUserID}", GetMessagesHandler)

	route.HandleFunc("/ws/{userID}", func(responseWriter http.ResponseWriter, request *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		// Reading username from request parameter
		userID := mux.Vars(request)["userID"]

		// Upgrading the HTTP connection socket connection
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		connection, err := upgrader.Upgrade(responseWriter, request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			log.Println("success")
		}
		CreateNewSocketUser(hub, connection, userID)

	})

	log.Println("Routes are Loaded.")

}
