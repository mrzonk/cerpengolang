package routes

import (
	"cerpengolang/controllers/CommentController"

	"github.com/gorilla/mux"
)

func CommentRoutes(r *mux.Router) {

	router := r.PathPrefix("/comment").Subrouter()

	router.HandleFunc("", CommentController.Index).Methods("GET")
	router.HandleFunc("", CommentController.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", CommentController.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", CommentController.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", CommentController.Delete).Methods("DELETE")
}
