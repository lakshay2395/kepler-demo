package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lakshay2395/kepler-demo/db"
	handler "github.com/lakshay2395/kepler-demo/handlers"
	"github.com/lakshay2395/kepler-demo/routes"
)

func initEnv() {
	os.Setenv("INFLUX_HOST", "localhost")
	os.Setenv("INFLUX_PORT", "8086")
	os.Setenv("INFLUX_USERNAME", "")
	os.Setenv("INFLUX_PASSWORD", "")
	os.Setenv("PORT", "8080")
}

func main() {
	initEnv() //TODO: to remove after finalization
	handler.CreateDB()
	err := db.Init(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_PORT"), os.Getenv("INFLUX_USERNAME"), os.Getenv("INFLUX_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}

// func main() {

// 	conClient := repository.CreateClient()
// 	repository.ExecuteQuery("create database stargate_ui", conClient)
// 	fmt.Println()
// }
