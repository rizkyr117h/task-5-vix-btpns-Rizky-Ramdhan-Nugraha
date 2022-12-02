package main

import (
	"log"
	"net/http"

	"github.com/rizkyr117h/middlewares"

	"github.com/rizkyr117h/controllers/authcontroller"
	"github.com/rizkyr117h/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")


	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
