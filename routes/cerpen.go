package routes

import (
	"cerpengolang/controllers/CerpenController"

	"github.com/gorilla/mux"
)

func CerpenRoutes(r *mux.Router) {

	router := r.PathPrefix("/cerpens").Subrouter()

	router.HandleFunc("", CerpenController.Index).Methods("GET")
	router.HandleFunc("", CerpenController.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", CerpenController.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", CerpenController.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", CerpenController.Delete).Methods("DELETE")
}
