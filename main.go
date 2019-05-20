package main

import (
	"go-playground/api"
	"go-playground/storage"
)

func main() {
	db := storage.CreateSqliteDeskStorage("test.db")
	defer db.Close()

	api.RunServer("localhost", 8080, db)
}
