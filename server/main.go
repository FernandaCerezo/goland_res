package main

import (
	"log"
	"net/http"

	"github.com/FernandaCerezo/golang_res/server/api"
	"github.com/gorilla/mux"
)

func main() {
	// db := database.GetConnection()
	//enrrutador
	router := mux.NewRouter().StrictSlash(true) //StrictSlash especifica la URL correcta
	apiRouter := router.PathPrefix("/api/").Subrouter()
	//Users
	apiRouter.HandleFunc("/users", api.GetAllUsers).Methods("GET")
	apiRouter.HandleFunc("/users", api.CreateUser).Methods("POST")
	//Reminders
	apiRouter.HandleFunc("/reminders", api.GetReminders).Methods("GET")
	apiRouter.HandleFunc("/reminders", api.CreateReminder).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
