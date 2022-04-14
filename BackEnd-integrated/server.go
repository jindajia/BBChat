package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	db "private-chat/database"
	helper "private-chat/helper"
	route2 "private-chat/route"
)

func main() {

	godotenv.Load()

	fmt.Println(
		fmt.Sprintf("%s%s%s%s", "Server will start at http://", os.Getenv("HOST"), ":", os.Getenv("PORT")),
	)

	db.ConnectDatabase()

	route := mux.NewRouter()

	route2.AddApproutes(route)

	serverPath := ":" + os.Getenv("PORT")

	cors := helper.GetCorsConfig()

	http.ListenAndServe(serverPath, cors.Handler(route))
}
