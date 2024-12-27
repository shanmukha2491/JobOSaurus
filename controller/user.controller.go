package controller

import (
	db "JobTracker/database"
	"JobTracker/models"
	"JobTracker/utils"
	"encoding/json"
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
