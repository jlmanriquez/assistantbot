package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jlmanriquez/assistantbot/assistantbot/dlgflow"

	"github.com/jlmanriquez/assistantbot/assistantbot/response"
)

// Talk recibe un POST con una estructura HumanResourceTalk en el body, que contiene el texto
// ingresado desde el cliente el cual es enviado a Dialogflow para detectar la intencion.
// La funcion retorna la respuesta obtenida desde Dialogflow.
func Talk(w http.ResponseWriter, req *http.Request) {
	var talk HumanResourceTalk

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&talk); err != nil {
		log.Printf(err.Error())
		resp.SendResponseBadReq(w, resp.Response{Code: "-1", Message: "Formato de entrada inválido"})
		return
	}

	result, err := dlgflow.SendIntent(talk.Text, talk.UserID)
	if err != nil {
		resp.SendResponseInternalErr(w, resp.Response{Code: "-1", Message: "No fue posible ejecutar la acción"})
		return
	}

	resp.SendResponse(w, result)
}
