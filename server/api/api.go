package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FernandaCerezo/golang_res/server/controllers"
	uuid "github.com/satori/go.uuid"
)

type Data_scheduled struct {
	Success bool                          `json:"success"`
	Data    []controllers.Scheduled_items `json:"data"`
	Errors  []string                      `json:"errors"`
}

type Data_users struct {
	Success bool                `json:"success"`
	Data    []controllers.Users `json:"data"`
	Errors  []string            `json:"errors"`
}

type response struct {
	ID      uuid.UUID `json:"UUID"`
	Message string    `json:"message,omitempty"`
}

func CreateReminder(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var reminder controllers.Scheduled_items

	err := json.NewDecoder(req.Body).Decode(&reminder)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := controllers.InsertReminder(reminder)
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetReminders(w http.ResponseWriter, req *http.Request) {
	var todos []controllers.Scheduled_items = controllers.GetAll()

	var data = Data_scheduled{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user controllers.Users

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := controllers.CreateUser(user)
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	var todos []controllers.Users = controllers.GetAllUsers()

	var data = Data_users{true, todos, make([]string, 0)}

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
