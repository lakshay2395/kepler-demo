package db

import (
	"fmt"
	"net/url"

	client "github.com/influxdata/influxdb1-client/v2"
)

var db *client.Client

func Init(hostname, port, username, password string) error {
	host, err := url.Parse(fmt.Sprintf("http://%s:%s", hostname, port))
	if err != nil {
		return err
	}
	conf := client.Config{
		URL:      *host,
		Username: username,
		Password: password,
	}
	c, err := client.NewClient(conf)
	if err != nil {
		return err
	}
	db = c
	return nil
}

func GetClient() *client.Client {
	return db
}
