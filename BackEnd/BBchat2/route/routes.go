package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	controllers "private-chat/controller"
)

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	log.Println("Loadeding Routes...")

	hub := controllers.NewHub()
	go hub.Run()

	route.HandleFunc("/", controllers.RenderHome)

	route.HandleFunc("/isUsernameAvailable/{username}", controllers.IsUsernameAvailable)

	route.HandleFunc("/login", controllers.Login).Methods("POST")

	route.HandleFunc("/registration", controllers.Registertation).Methods("POST")

	route.HandleFunc("/userSessionCheck/{userID}", controllers.UserSessionCheck)

	route.HandleFunc("/getConversation/{toUserID}/{fromUserID}", controllers.GetMessagesHandler)

    route.HandleFunc("/getBroadcast/{fromUserID}", controllers.GetBroadcastHandler)

    route.HandleFunc("/getDriftBottle/{toUserID}/{fromUserID}", controllers.GetDriftBottlesHandler)

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
		controllers.CreateNewSocketUser(hub, connection, userID)

	})

	log.Println("Routes are Loaded.")
}
