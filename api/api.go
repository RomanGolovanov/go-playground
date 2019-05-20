package api

import (
	"encoding/json"
	"go-playground/model"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func getDesk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("id")
	if name == "" {
		http.Error(w, "Invalid desk name", http.StatusBadRequest)
		return
	}
	desk, ok := model.GetDesk(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desk)
}

func deleteDesk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("id")
	if name == "" {
		http.Error(w, "Invalid desk name", http.StatusBadRequest)
		return
	}
	ok := model.DeleteDesk(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.StatusText(http.StatusNoContent)
}

func getDeskIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	desks := model.GetAllDesks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desks)
}

func createOrUpdateDesk(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Body == nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	var desk model.Desk
	err := json.NewDecoder(r.Body).Decode(&desk)
	if err != nil {
		http.Error(w, "Invalid body format", http.StatusBadRequest)
		return
	}
	model.NewDesk(desk)
	http.StatusText(http.StatusNoContent)
}

// RunServer starts api listener
func RunServer(hostName string, port int) {
	router := httprouter.New()

	router.GET("/api/1.0/desk/:id", getDesk)
	router.DELETE("/api/1.0/desk/:id", deleteDesk)
	router.POST("/api/1.0/desk", createOrUpdateDesk)
	router.GET("/api/1.0/desk", getDeskIndex)

	address := hostName + ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(address, router))
}
