package handlers

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

type MapBox struct {
	Type     string          `json:"type"`
	Features []MapBoxFeature `json:"features"`
}

type MapBoxFeature struct {
	Type       string                `json:"type"`
	Properties MapBoxFeatureProperty `json:"properties"`
	Geometry   MapBoxGeometry        `json:"geometry"`
}

type MapBoxFeatureProperty struct {
	DBH interface{} `json:"dbh"`
}

type MapBoxGeometry struct {
	Type        string        `json:"point"`
	Coordinates []json.Number `json:"coordinates"`
}

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
	enableCors(&w)
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
	enableCors(&w)
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
	enableCors(&w)
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
		query = generateSupplyCommands("select * from low_supply ", serviceArea, serviceType)
		break
	case RAIN_CHECK:
		query = generateSupplyCommands("select * from rain_check ", serviceArea, serviceType)
		break
	}
	rows := getResponse(query)[0].Series[0].Values
	payload := []MapBoxFeature{}
	for _, row := range rows {
		feature := MapBoxFeature{}
		feature.Type = "Feature"
		feature.Properties = MapBoxFeatureProperty{DBH: row[len(row)-1]}
		feature.Geometry = MapBoxGeometry{}
		feature.Geometry.Type = "Point"
		feature.Geometry.Coordinates = []json.Number{row[1].(json.Number), row[2].(json.Number)}
		payload = append(payload, feature)
	}
	response := MapBox{}
	response.Type = "FeatureCollection"
	response.Features = payload
	Ok(w, response)
}

func generateSupplyCommands(service string, serviceArea string, serviceType string) string {
	var query string
	if len(serviceArea) > 0 && len(serviceType) > 0 {
		query = service + "where service_area='" + serviceArea + "' and service_type='" + serviceType + "'"
	} else if len(serviceArea) > 0 {
		query = service + "where service_area='" + serviceArea + "'"
	} else if len(serviceType) > 0 {
		query = service + "where service_type='" + serviceType + "'"
	} else {
		query = service
	}
	fmt.Println(query)
	return query
}

func GetServiceAreas(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
