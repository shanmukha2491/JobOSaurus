package routes

import (
	handler "JobTracker/controller"
	"JobTracker/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterUserRoutes sets up the routes for user-related actions
func RegisterUserRoutes(router *mux.Router) {

	// Test route to check if the server is running
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.NewApiResponse(200, "Test", "Server is running"))
	}).Methods(http.MethodGet) // It's better to use GET for health check endpoints

	// Route for creating a user (POST request)
	router.HandleFunc("/v1/user", handler.CreateUser).Methods(http.MethodPost)

	router.HandleFunc("/v1/user/login", handler.LoginUser).Methods(http.MethodPost)
}
