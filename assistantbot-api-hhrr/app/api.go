package app

import (
	"log"
	"net/http"

	"github.com/jlmanriquez/assistantbot-api-hhrr/app/response"

	"github.com/gorilla/mux"
)

// GetInsurances retorna la lista de seguros contratados del usuario.
func GetInsurances(w http.ResponseWriter, req *http.Request) {
	userID := mux.Vars(req)["userID"]

	log.Printf("Consultando seguros contratados para usuario %s", userID)

	resp.SendResponse(w, []string{"Hogar", "Vida"})
}
