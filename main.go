package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type reminder struct {
	CreatedBy int           `json:"user_id"`
	Title     string        `json:"title"`
	Message   string        `json:"message"`
	Duration  time.Duration `json:"duration"`
}

type allReminders []reminder

var reminders = allReminders{
	{
		CreatedBy: 1,
		Title:     "Task One",
		Message:   "Some Content",
		Duration:  3,
	},
}

func getReminders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reminders)
}

func createReminder(w http.ResponseWriter, r *http.Request) {
	var newReminder reminder
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "insert correct values")

	}

	json.Unmarshal(reqBody, &newReminder)
	newReminder.CreatedBy = len(reminders) + 1
	reminders = append(reminders, newReminder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newReminder)
}

func deleteReminder(w http.ResponseWriter, r *http.Request) {
	//Extraer la variable desde la peticion
	vars := mux.Vars(r)
	reminderId, err := strconv.Atoi(vars["user_id"])

	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}
	for index, item := range reminders {
		if item.CreatedBy == reminderId {
			reminders = append(reminders[:index], reminders[index+1:]...)
			fmt.Fprintf(w, "The reminder with ID %v has been remove succesfully", reminderId)
		}
	}

}

func updateReminder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reminderId, err := strconv.Atoi(vars["user_id"])
	var updatedReminder reminder
	if err != nil {
		fmt.Fprintf(w, "invalid ID")
		return
	}

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
		return
	}

	json.Unmarshal(reqBody, &updatedReminder)
	for index, item := range reminders {
		if item.CreatedBy == reminderId {
			reminders = append(reminders[:index], reminders[index+1:]...)
			updatedReminder.CreatedBy = reminderId
			reminders = append(reminders, updatedReminder)

			fmt.Fprintf(w, "The reminder with ID %v has been updated succesfully", reminderId)

		}
	}
}

func main() {
	//enrrutador
	router := mux.NewRouter().StrictSlash(true) //StrictSlash especifica la URL correcta
	router.HandleFunc("/", getReminders).Methods("GET")
	router.HandleFunc("/create", createReminder).Methods("POST")
	router.HandleFunc("/reminder/{CreatedBy}", deleteReminder).Methods("DELETE")
	router.HandleFunc("/reminder/{CreatedBy}", updateReminder).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
