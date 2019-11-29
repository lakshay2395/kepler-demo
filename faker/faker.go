package faker

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client"
	"github.com/lakshay2395/kepler-demo/db"
	handler "github.com/lakshay2395/kepler-demo/handlers"
)

func GenerateDataFromCSV() {
	csv_file, ferr := os.Open(os.Getenv("CSV_FILE"))
	if ferr != nil {
		log.Fatal(ferr)
	}
	r := csv.NewReader(csv_file)
	pts := []client.Point{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		latValue, err := strconv.ParseFloat(record[6], 32)
		longValue, err := strconv.ParseFloat(record[7], 32)

		pts = append(pts, client.Point{
			Measurement: db.LOW_SUPPLY_MEASUREMENT,
			Tags: map[string]string{
				"service_area": "Jabodetabek",
				"service_type": "car",
			},
			Fields: map[string]interface{}{
				"lat":   latValue,
				"long":  longValue,
				"value": rand.Float64(),
			},
			Time: time.Now(),
		})
	}
	writeDataToDB(pts)
}

func writeDataToDB(pts []client.Point) {

	c := db.GetClient()
	bps := client.BatchPoints{
		Points:   pts,
		Database: db.GetDBName(),
	}
	_, err := c.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

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
