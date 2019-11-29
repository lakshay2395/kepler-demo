package main

import (
	"fmt"
	"os"

	"github.com/lakshay2395/kepler-demo/repository"
)

func initEnv() {
	os.Setenv("INFLUX_HOST", "localhost")
	os.Setenv("INFLUX_PORT", "8086")
	os.Setenv("INFLUX_USERNAME", "")
	os.Setenv("INFLUX_PASSWORD", "")
	os.Setenv("PORT", "8080")
}

// func main() {
// 	initEnv() //TODO: to remove after finalization
// 	err := db.Init(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_PORT"), os.Getenv("INFLUX_USERNAME"), os.Getenv("INFLUX_PASSWORD"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r := mux.NewRouter()
// 	routes.RegisterRoutes(r)
// 	http.Handle("/", r)
// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
// }

func main() {

	conClient := repository.CreateClient()
	repository.ExecuteQuery("create database stargate_ui", conClient)
	fmt.Println()
}
