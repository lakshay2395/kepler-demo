package db

import (
	"fmt"
	"net/url"
	"os"

	client "github.com/influxdata/influxdb1-client"
)

var db *client.Client
var RAIN_CHECK_MEASUREMENT = "rain_check"
var LOW_SUPPLY_MEASUREMENT = "low_supply"

func Init(hostname, port, username, password, dbName string) error {
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

func CreateDB(dbName string) error {
	response, err := db.Query(client.Query{
		Command:  "create database " + dbName,
		Database: dbName,
	})
	if err != nil {
		return err
	}
	if response.Error() != nil {
		return response.Error()
	}
	return nil
}

func DropDB(dbName string) error {
	response, err := db.Query(client.Query{
		Command:  "drop database " + dbName,
		Database: dbName,
	})
	if err != nil {
		return err
	}
	if response.Error() != nil {
		return response.Error()
	}
	return nil
}

func GetClient() *client.Client {
	return db
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}
