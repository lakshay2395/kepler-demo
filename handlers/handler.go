package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/influxdata/influxdb1-client"
	"github.com/lakshay2395/kepler-demo/db"
)

//Default db name
var DB_NAME string = "BumbeBeeTuna"

func GetTripsList(w http.ResponseWriter, r *http.Request) {
	executeQuery(w, "select * from shapes")
}

func Error(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	response := map[string]string{}
	response["error"] = fmt.Sprintf("%s", err)
	payload, _ := json.Marshal(response)
	w.Write(payload)
}

func executeQuery(w http.ResponseWriter, query string) {
	db := db.GetClient()
	response, err := db.Query(client.Query{
		Command:  query,
		Database: DB_NAME,
	})
	if err != nil {
		Error(w, err)
		return
	}
	if response.Error() != nil {
		Error(w, response.Error())
		return
	}
	Ok(w, response.Results)
}

func Ok(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	payload, _ := json.Marshal(response)
	w.Write(payload)
}
