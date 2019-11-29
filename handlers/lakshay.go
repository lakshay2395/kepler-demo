package handler

import "net/http"

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
	Ok(w, data)
}

func GetMetricRecommendations(w http.ResponseWriter, r *http.Request) {

	// data, err := ReadFile("metric_recommendations")
	// if err != nil {
	// 	Error(w, err)
	// 	return
	// }
	// recommendations := []MetricRecommendation{}
	// err = json.Unmarshal(data, &payload)
	// if err != nil {
	// 	Error(w, err)
	// 	return
	// }
	// rec := []MetricRecommendation{}
	// for _,recommendation := range recommendations{
	// 	if recommendation.ID
	// }
	// Ok(w, data)
}
