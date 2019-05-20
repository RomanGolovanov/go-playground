package main

import (
	"go-playground/api"
	"go-playground/storage"
)

func main() {
	api.RunServer("localhost", 8080, storage.CreateMemoryDeskStorage())
}
