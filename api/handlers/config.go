package handlers

import "github.com/gorilla/mux"


func ConfigureRouter(handler *Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.home).Methods("GET")
	r.HandleFunc("/products", handler.getProducts).Methods("GET")
	return r
}