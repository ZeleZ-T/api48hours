package noiseMap

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetRoutes(r *chi.Mux) {
	r.Post("/world-map", createMap)
	r.Get("/world-map", getMap)
	r.Patch("/world-map", updateMapName)
	r.Delete("/world-map", deleteMap)
}

func createMap(w http.ResponseWriter, r *http.Request) {

}

func getMap(w http.ResponseWriter, r *http.Request) {

}

func updateMapName(w http.ResponseWriter, r *http.Request) {

}

func deleteMap(w http.ResponseWriter, r *http.Request) {

}
