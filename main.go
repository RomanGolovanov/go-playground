package main

import (
	"go-playground/api"
	"go-playground/storage"
)

func main() {
	db := storage.CreateSqliteDeskStorage("test.sqlite")
	defer db.Close()

	api.RunServer("localhost", 8080, db)
}
