package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jlmanriquez/assistantbot-api-hhrr/app"
	"github.com/jlmanriquez/assistantbot-api-hhrr/app/response"
)

func main() {
	r := configureRouter()

	log.Fatal(http.ListenAndServe(":8001", r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods")
		next.ServeHTTP(w, r)
	})
}

func manageCoorsOptions(w http.ResponseWriter, r *http.Request) {
	resp.SendResponse(w, nil)
}

func configureRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)

	// Configuracion para manejar las peticiones OPTIONS
	r.Methods("OPTIONS").HandlerFunc(manageCoorsOptions)

	// Se agregan los endpoints de los modulos
	app.AddRoutes(r)

	return r
}
