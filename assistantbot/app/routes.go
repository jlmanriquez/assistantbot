package app

import "github.com/gorilla/mux"

// AddRoutes configura los endpoints del modulo
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/app/talk", Talk).Methods("POST")
}
