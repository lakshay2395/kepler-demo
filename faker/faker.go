package faker

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	client "github.com/influxdata/influxdb1-client"
	"github.com/lakshay2395/kepler-demo/db"
	handler "github.com/lakshay2395/kepler-demo/handlers"
)

func GenerateLowSupplyData() {
	c := db.GetClient()
	serviceAreas, err := GetServiceAreas()
	if err != nil {
		log.Fatal(err)
	}
	serviceTypes, err := GetServiceTypes()
	if err != nil {
		log.Fatal(err)
	}
	pts := []client.Point{}
	for _, serviceArea := range serviceAreas {
		if IsLatLongPresentInServiceArea(serviceArea) {
			for _, serviceType := range serviceTypes {
				pts = append(pts, GenerateClientPoint(db.LOW_SUPPLY_MEASUREMENT, serviceArea, serviceType, rand.Float64(), time.Now()))
			}
		}
	}
	bps := client.BatchPoints{
		Points:   pts,
		Database: db.GetDBName(),
	}
	_, err = c.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateRainCheckData() {
	c := db.GetClient()
	serviceAreas, err := GetServiceAreas()
	if err != nil {
		log.Fatal(err)
	}
	serviceTypes, err := GetServiceTypes()
	if err != nil {
		log.Fatal(err)
	}
	pts := []client.Point{}
	for _, serviceArea := range serviceAreas {
		if IsLatLongPresentInServiceArea(serviceArea) {
			for _, serviceType := range serviceTypes {
				pts = append(pts, GenerateClientPoint(db.RAIN_CHECK_MEASUREMENT, serviceArea, serviceType, rand.Float64(), time.Now()))
			}
		}
	}
	bps := client.BatchPoints{
		Points:   pts,
		Database: db.GetDBName(),
	}
	_, err = c.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func IsLatLongPresentInServiceArea(serviceArea handler.ServiceArea) bool {
	return serviceArea.Lat != 0 && serviceArea.Long != 0
}

func GetServiceAreas() ([]handler.ServiceArea, error) {
	data, err := handler.ReadFile("data/service_areas")
	if err != nil {
		return nil, err
	}
	serviceAreas := []handler.ServiceArea{}
	err = json.Unmarshal(data, &serviceAreas)
	if err != nil {
		return nil, err
	}
	return serviceAreas, nil
}

func GetServiceTypes() ([]handler.ServiceType, error) {
	data, err := handler.ReadFile("data/service_types")
	if err != nil {
		return nil, err
	}
	serviceTypes := []handler.ServiceType{}
	err = json.Unmarshal(data, &serviceTypes)
	if err != nil {
		return nil, err
	}
	return serviceTypes, nil
}

func GenerateClientPoint(measurementName string, serviceArea handler.ServiceArea, serviceType handler.ServiceType, value interface{}, t time.Time) client.Point {
	return client.Point{
		Measurement: measurementName,
		Tags: map[string]string{
			"service_area": serviceArea.Name,
			"service_type": serviceType.Name,
		},
		Fields: map[string]interface{}{
			"lat":   serviceArea.Lat,
			"long":  serviceArea.Long,
			"value": value,
		},
		Time: t,
	}
}
