package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	LOW_SUPPLY = "1"
	RAIN_CHECK = "2"
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

type ServiceArea struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type ServiceType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	vars := mux.Vars(r)
	var serviceType, serviceArea string
	_, ok := r.URL.Query()["serviceArea"]
	if ok {
		serviceArea = r.URL.Query()["serviceArea"][0]
	}
	_, ok = r.URL.Query()["serviceType"]
	if ok {
		serviceType = r.URL.Query()["serviceType"][0]
	}
	var query string
	switch vars["id"] {
	case LOW_SUPPLY:
		query = generateSupplyCommands("select * from low_supply where ", serviceArea, serviceType)
		break
	case RAIN_CHECK:
		query = generateSupplyCommands("select * from rain_check where ", serviceArea, serviceType)
		break
	}
	response := getResponse(query)[0]
	columns := response.Series[0].Columns
	values := response.Series[0].Values

	data := createMap(columns, values)

	Ok(w, data)
}

func generateSupplyCommands(service string, serviceArea string, serviceType string) string {
	var query string
	if len(serviceArea) > 0 && len(serviceType) > 0 {
		query = service + "service_area='" + serviceArea + "' and service_type='" + serviceType + "'"
	} else if len(serviceArea) > 0 {
		query = service + "service_area='" + serviceArea + "'"
	} else if len(serviceType) > 0 {
		query = service + "service_types='" + serviceType + "'"
	}
	fmt.Println(query)
	return query
}

func GetServiceAreas(w http.ResponseWriter, r *http.Request) {
	data, err := ReadFile("service_areas")
	if err != nil {
		Error(w, err)
		return
	}
	serviceAreas := []ServiceArea{}
	err = json.Unmarshal(data, &serviceAreas)
	if err != nil {
		Error(w, err)
		return
	}
	Ok(w, serviceAreas)
}

func GetServiceTypes(w http.ResponseWriter, r *http.Request) {
	data, err := ReadFile("service_types")
	if err != nil {
		Error(w, err)
		return
	}
	serviceTypes := []ServiceType{}
	err = json.Unmarshal(data, &serviceTypes)
	if err != nil {
		Error(w, err)
		return
	}
	Ok(w, serviceTypes)
}

func createMap(columns []string, values [][]interface{}) map[string]interface{} {
	date := make(map[string]interface{})

	for i := 0; i < len(columns); i++ {
		fmt.Println(columns[i], values[i])
		data[columns[i]] = values[i]
	}
	return data
}
