package routes

import "github.com/gorilla/mux"

import handler "github.com/lakshay2395/kepler-demo/handlers"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/trips", handler.GetTripsList).Methods("GET")
}
