package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Metric struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MetricRecommendation struct {
	ID       string `json:"id"`
	MetricID string `json:"metric_id"`
	Name     string `json:"name"`
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	data, err := ReadFile("metric_types")
	if err != nil {
		Error(w, err)
		return
	}
	metrics := []Metric{}
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		Error(w, err)
		return
	}
	Ok(w, metrics)
}

func GetMetricRecommendations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := ReadFile("metric_recommendations")
	if err != nil {
		Error(w, err)
		return
	}
	recommendations := []MetricRecommendation{}
	err = json.Unmarshal(data, &recommendations)
	if err != nil {
		Error(w, err)
		return
	}
	list := []MetricRecommendation{}
	for _, recommendation := range recommendations {
		if recommendation.MetricID == vars["id"] {
			list = append(list, recommendation)
		}
	}
	Ok(w, list)
}

func GetMetricsData(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	executeQuery(w, "select * from shapes")
}
