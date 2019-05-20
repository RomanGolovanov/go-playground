package api

import (
	"encoding/json"
	model "github.com/RomanGolovanov/go-playground/model"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func getDeskIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	desks := model.GetAllDesks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desks)
}

func getDesk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	desk, ok := model.GetDesk(ps.ByName("id"))

	if !ok {
		http.Error(w, "Not Found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desk)
}

// RunServer starts api listener
func RunServer(hostName string, port int) {
	router := httprouter.New()

	router.GET("/api/1.0/desk/:id", getDesk)
	router.GET("/api/1.0/desk", getDeskIndex)

	address := hostName + ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(address, router))
}
