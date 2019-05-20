package api

import (
	"encoding/json"
	"go-playground/model"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// IDeskStorage should be implemented to process api requests
type IDeskStorage interface {
	NewDesk(desk model.Desk)
	GetDesk(name string) (model.Desk, bool)
	DeleteDesk(name string) bool
	GetAllDesks() []model.Desk
}

func getDesk(w http.ResponseWriter, r *http.Request, ps httprouter.Params, storage IDeskStorage) {
	name := ps.ByName("id")
	if name == "" {
		http.Error(w, "Invalid desk name", http.StatusBadRequest)
		return
	}
	desk, ok := storage.GetDesk(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desk)
}

func deleteDesk(w http.ResponseWriter, r *http.Request, ps httprouter.Params, storage IDeskStorage) {
	name := ps.ByName("id")
	if name == "" {
		http.Error(w, "Invalid desk name", http.StatusBadRequest)
		return
	}
	ok := storage.DeleteDesk(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.StatusText(http.StatusNoContent)
}

func getDeskIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params, storage IDeskStorage) {
	desks := storage.GetAllDesks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(desks)
}

func createOrUpdateDesk(w http.ResponseWriter, r *http.Request, _ httprouter.Params, storage IDeskStorage) {
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
	storage.NewDesk(desk)
	http.StatusText(http.StatusNoContent)
}

// RunServer starts api listener
func RunServer(hostName string, port int, storage IDeskStorage) {
	router := httprouter.New()

	router.GET("/api/1.0/desk/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		getDesk(w, r, ps, storage)
	})
	router.DELETE("/api/1.0/desk/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		deleteDesk(w, r, ps, storage)
	})
	router.POST("/api/1.0/desk", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		createOrUpdateDesk(w, r, ps, storage)
	})
	router.GET("/api/1.0/desk", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		getDeskIndex(w, r, ps, storage)
	})

	address := hostName + ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(address, router))
}
