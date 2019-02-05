package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/jlmanriquez/assistantbot/assistantbot/app"
	"github.com/jlmanriquez/assistantbot/assistantbot/response"
)

func main() {
	runCredentials()

	// Se ejecuta cuando la aplicacion sea finalizada mediante Ctrl-c
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Println("Shutdown...")
		os.Exit(0)
	}()

	r := configureRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}

func runCredentials() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filename := os.Args[1]

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", dir+filename)
	log.Printf("Credenciales ejecutadas correctamente: %s", dir+filename)
}

func addCOORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods")
		next.ServeHTTP(w, r)
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func manageOptionsCOORS(w http.ResponseWriter, r *http.Request) {
	resp.SendResponse(w, nil)
}

func configureRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(addCOORSMiddleware)
	r.Use(logMiddleware)

	// Configuracion para manejar las peticiones OPTIONS
	r.Methods("OPTIONS").HandlerFunc(manageOptionsCOORS)

	// Se agregan los endpoints de los modulos
	app.AddRoutes(r)

	return r
}
