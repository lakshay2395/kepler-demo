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
	data, err := handler.ReadFile("service_areas")
	if err != nil {
		log.Fatal(err)
	}
	pts := []client.Point{}
	serviceAreas := []handler.ServiceArea{}
	err = json.Unmarshal(data, &serviceAreas)
	if err != nil {
		log.Fatal(err)
	}
	data, err = handler.ReadFile("service_types")
	if err != nil {
		log.Fatal(err)
	}
	serviceTypes := []handler.ServiceType{}
	err = json.Unmarshal(data, &serviceTypes)
	if err != nil {
		log.Fatal(err)
	}
	for _, serviceArea := range serviceAreas {
		if serviceArea.Lat == 0 || serviceArea.Long == 0 {
			break
		}
		for _, serviceType := range serviceTypes {
			pts = append(pts, client.Point{
				Measurement: db.LOW_SUPPLY_MEASUREMENT,
				Tags: map[string]string{
					"service_area": serviceArea.Name,
					"service_type": serviceType.Name,
				},
				Fields: map[string]interface{}{
					"lat":   serviceArea.Lat,
					"long":  serviceArea.Long,
					"value": randFloats(0.0, 1.0, 1)[0],
				},
				Time: time.Now(),
			})
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
	data, err := handler.ReadFile("service_areas")
	if err != nil {
		log.Fatal(err)
	}
	pts := []client.Point{}
	serviceAreas := []handler.ServiceArea{}
	err = json.Unmarshal(data, &serviceAreas)
	if err != nil {
		log.Fatal(err)
	}
	data, err = handler.ReadFile("service_types")
	if err != nil {
		log.Fatal(err)
	}
	serviceTypes := []handler.ServiceType{}
	err = json.Unmarshal(data, &serviceTypes)
	if err != nil {
		log.Fatal(err)
	}
	for _, serviceArea := range serviceAreas {
		if serviceArea.Lat == 0 || serviceArea.Long == 0 {
			break
		}
		for _, serviceType := range serviceTypes {
			pts = append(pts, client.Point{
				Measurement: db.RAIN_CHECK_MEASUREMENT,
				Tags: map[string]string{
					"service_area": serviceArea.Name,
					"service_type": serviceType.Name,
				},
				Fields: map[string]interface{}{
					"lat":   serviceArea.Lat,
					"long":  serviceArea.Long,
					"value": randBoolean(randFloats(0.0, 1.0, 1)[0]),
				},
				Time: time.Now(),
			})
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

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func randBoolean(x float64) bool {
	if x > 0.5 {
		return true
	}
	return false
}
