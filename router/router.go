package router

import (
	"main/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/movie", controller.Insert).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkeAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAll).Methods("DELETE")

	return router
}
