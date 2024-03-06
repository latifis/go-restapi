package main

import (
	"log"
	"net/http"

	"go-restapi/controllers/authcontroller"
	"go-restapi/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
