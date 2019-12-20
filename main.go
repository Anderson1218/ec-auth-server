package main

import (
	"database/sql"
	"github.com/Anderson1218/ec-auth-server/controllers"
	"github.com/Anderson1218/ec-auth-server/driver"
	"github.com/Anderson1218/ec-auth-server/utils"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/api/users/", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/api/users/token", controller.Login(db)).Methods("POST")
	router.HandleFunc("/api/users/me", utils.TokenVerifyMiddleWare(controller.Me(db))).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
