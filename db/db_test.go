package db

import (
	"os"
	"testing"
)

func TestGetClient(t *testing.T) {
	hostname := "localhost"
	dbName := "keplerDemo"
	port := "8090"
	Init(hostname, port, "", "", "")
	dB := CreateDB(dbName)
	if dB == nil {
		t.Errorf("dB creation failed due to some error")
	}
}

func TestGetDBName(t *testing.T) {
	dbName := GetDBName()
	if dbName != os.Getenv("DB_NAME") {
		t.Fatalf("Expected value to be %v, but got %v", dbName, os.Getenv("DB_NAME"))
	}
}

func TestDropDB(t *testing.T) {
	dBName := GetDBName()
	DropDB(dBName)
	if GetClient() != nil {
		t.Error("Error delete database !")
	}
	hostname := "localhost"
	port := "8090"
	Init(hostname, port, "", "", "")
}
