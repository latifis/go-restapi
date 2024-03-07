package main

import (
	"log"
	"net/http"

	"go-restapi/middlewares"

	"go-restapi/controllers/authcontroller"
	"go-restapi/controllers/karyawancontroller"
	"go-restapi/models"

	"github.com/gorilla/mux"
)

func main() {

	// Menghubungkan ke database
	models.ConnectDatabase()

	// Membuat router baru
	r := mux.NewRouter()

	// Rute untuk autentikasi
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	// Rute untuk karyawan dengan autentikasi JWT
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/karyawan", karyawancontroller.GetKaryawan).Methods("GET")
	api.HandleFunc("/karyawan/{id}", karyawancontroller.GetKaryawanByID).Methods("GET")
	api.HandleFunc("/karyawan", karyawancontroller.CreateKaryawan).Methods("POST")
	api.HandleFunc("/karyawan/{id}", karyawancontroller.UpdateKaryawan).Methods("PUT")
	api.HandleFunc("/karyawan/{id}", karyawancontroller.DeleteKaryawan).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	// Menjalankan server HTTP
	log.Fatal(http.ListenAndServe(":8080", r))
}
