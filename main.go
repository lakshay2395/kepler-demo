package main

import (
	"fmt"
	repo "github.com/lakshay2395/kepler-demo/repository"
)

func main() {

	conClient := repo.CreateClient()
	repo.ExecuteQuery("create database stargate_ui", conClient)
	fmt.Println()
}
