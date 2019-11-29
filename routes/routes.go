package routes

import "github.com/gorilla/mux"

import handler "github.com/lakshay2395/kepler-demo/handlers"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/metrics", handler.GetMetrics).Methods("GET")
	router.HandleFunc("/metrics/{id}/recommendations", handler.GetMetricRecommendations).Methods("GET")
	router.HandleFunc("/metrics/{id}/data", handler.GetMetricRecommendations).Methods("GET")
}
