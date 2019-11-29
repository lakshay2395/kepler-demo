package routes

import "github.com/gorilla/mux"

import handler "github.com/lakshay2395/kepler-demo/handlers"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/metrics", handler.GetMetrics).Methods("GET")
	router.HandleFunc("/metrics/{id}/recommendations", handler.GetMetricRecommendations).Methods("GET")
	router.HandleFunc("/metrics/{id}/data", handler.GetMetricsData).Methods("GET")
	router.HandleFunc("/service-types", handler.GetServiceTypes).Methods("GET")
	router.HandleFunc("/service-areas", handler.GetServiceAreas).Methods("GET")
}
