package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	client "github.com/influxdata/influxdb1-client"
	"github.com/lakshay2395/kepler-demo/db"
)

func Error(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	response := map[string]string{}
	response["error"] = fmt.Sprintf("%s", err)
	payload, _ := json.Marshal(response)
	w.Write(payload)
}

func executeQuery(w http.ResponseWriter, query string) []client.Result {
	c := db.GetClient()
	response, err := c.Query(client.Query{
		Command:  query,
		Database: db.GetDBName(),
	})
	if err != nil {
		Error(w, err)
		return nil
	}
	if response.Error() != nil {
		Error(w, response.Error())
		return nil
	}
	Ok(w, response.Results)
	return response.Results
}

func getResponse(query string) []client.Result {
	c := db.GetClient()
	response, err := c.Query(client.Query{
		Command:  query,
		Database: db.GetDBName(),
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return response.Results
}

func Ok(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	payload, _ := json.Marshal(response)
	w.Write(payload)
}

func ReadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(fmt.Sprintf("%s.json", name))
}
