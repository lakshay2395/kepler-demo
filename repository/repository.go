package repository

import (
	"fmt"
	"log"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

//Default db name
var DB_NAME string = "stargate_ui"

//Create connection client
func CreateClient() client.Client {
	host := fmt.Sprintf("http://%s:%d", "localhost", 8086)

	con, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: host,
	})
	if err != nil {
		log.Fatal(err)
	}
	return con
}

// Execute Influx Query
func ExecuteQuery(query string, conClient client.Client) {
	q := client.NewQuery(query, DB_NAME, "")
	if response, err := conClient.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}
