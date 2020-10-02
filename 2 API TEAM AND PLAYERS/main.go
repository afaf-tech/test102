package main

import (
	"fmt"
	"net/http"

	"./configuration"
	"./handlers"

	"github.com/gorilla/mux"
)

var CF, _ = configuration.LoadConfiguration()

func main() {
	fmt.Println("Starting the application")
	fmt.Println("Running On Port:", CF.Port)

	router := mux.NewRouter()
	// person
	router.HandleFunc("/players", handlers.GetPlayerEndpoint).Methods("GET")
	router.HandleFunc("/player", handlers.CreatePlayerEndpoint).Methods("POST")
	//team
	router.HandleFunc("/teams", handlers.GetTeamEndpoint).Methods("GET")
	router.HandleFunc("/team", handlers.CreateTeamEndpoint).Methods("POST")
	http.ListenAndServe(CF.Port, router)
}
