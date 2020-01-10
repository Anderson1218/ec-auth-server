package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Anderson1218/ec-auth-server/controllers"
	"github.com/Anderson1218/ec-auth-server/driver"
	"github.com/Anderson1218/ec-auth-server/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		// log.Fatal("$PORT must be set")
	}
	db = driver.ConnectDB()
	r := mux.NewRouter()

	controller := controllers.Controller{}

	r.HandleFunc("/users", controller.Signup(db)).Methods("POST")
	r.HandleFunc("/users/token", controller.Login(db)).Methods("POST")
	r.HandleFunc("/users/me", utils.TokenVerifyMiddleWare(controller.Me(db))).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
