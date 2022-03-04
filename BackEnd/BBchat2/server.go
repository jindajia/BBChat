package main

import (
	"fmt"
	"net/http"
	"os"
	routes "private-chat/route"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	db "private-chat/database"
	helper "private-chat/helper"
)

func main() {

	godotenv.Load()

	fmt.Println(
		fmt.Sprintf("%s%s%s%s", "Server will start at http://", os.Getenv("HOST"), ":", os.Getenv("PORT")),
	)

	db.ConnectDatabase()

	route := mux.NewRouter()

	routes.AddApproutes(route)

	serverPath := ":" + os.Getenv("PORT")

	cors := helper.GetCorsConfig()

	http.ListenAndServe(serverPath, cors.Handler(route))
}
