package main

import (
	"go-playground/api"
	"go-playground/storage"
	"fmt"
)

func main() {
	fmt.Println("Connecting to database.")

	db := storage.CreateSqliteDeskStorage("test.sqlite")
	defer db.Close()

	fmt.Println("Running web service.")

	api.RunServer("", 8080, db)
}
