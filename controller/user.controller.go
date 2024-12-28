package controller

import (
	db "JobTracker/database"
	"JobTracker/models"
	"JobTracker/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(rw, "Invalid user details", http.StatusBadRequest)
		return
	}

	err = db.InsertOne(newUser)
	if err != nil {
		http.Error(rw, "Error creating user", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(rw).Encode(utils.NewApiResponse(201, "Successfully Created User", newUser))
	if err != nil {
		http.Error(rw, "Internal Server Problem", http.StatusInternalServerError)
		return
	}

}

type loginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(rw http.ResponseWriter, r *http.Request) {
	var loginCredentials loginDetails
	err := json.NewDecoder(r.Body).Decode(&loginCredentials)

	if loginCredentials.Password == "" || loginCredentials.Username == "" {
		http.Error(rw, "invalid data", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "invalid data", http.StatusBadRequest)
		return
	}

	log.Print("loginCredentials: ", loginCredentials)
	user, err := db.FindOne(loginCredentials.Username, loginCredentials.Password)
	if err != nil {
		http.Error(rw, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(rw).Encode(utils.NewApiResponse(200, "Successfully Created User", user))
	if err != nil {
		http.Error(rw, "Internal Server Problem", http.StatusInternalServerError)
		return
	}
}
