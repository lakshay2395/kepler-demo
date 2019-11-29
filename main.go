package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lakshay2395/kepler-demo/db"
	"github.com/lakshay2395/kepler-demo/faker"
	"github.com/lakshay2395/kepler-demo/routes"
)

func initEnv() {
	os.Setenv("INFLUX_HOST", "localhost")
	os.Setenv("INFLUX_PORT", "8086")
	os.Setenv("INFLUX_USERNAME", "")
	os.Setenv("INFLUX_PASSWORD", "")
	os.Setenv("PORT", "8090")
	os.Setenv("DB_NAME", "keplerDemo")
	os.Setenv("LOW_SUPPLY_DATA", "data/Jakarta_data.csv")
	os.Setenv("RAIN_CHECK_DATA", "data/sample-rain-check.csv")
}

func main() {
	InitDB()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}

func InitDB() error {
	initEnv() //TODO: to remove after finalization
	dbName := os.Getenv("DB_NAME")
	err := db.Init(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_PORT"), os.Getenv("INFLUX_USERNAME"), os.Getenv("INFLUX_PASSWORD"), dbName)
	if err != nil {
		return err
	}
	err = db.DropDB(dbName)
	if err != nil {
		return err
	}
	err = db.CreateDB(dbName)
	if err != nil {
		return err
	}
	faker.GenerateLowSupplyData()
	faker.GenerateRainCheckData()
	faker.GenerateDataFromCSVForLowSupply()
	faker.GenerateDataFromCSVForRainCheck()
	return nil
}
