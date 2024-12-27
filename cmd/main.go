package main

import (
	"JobTracker/database"
	"JobTracker/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Job Application Tracker")
	
	database.ConnectMongo()

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
