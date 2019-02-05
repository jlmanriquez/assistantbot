package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response es la estructura de salida que retorna la aplicacion
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// SendResponse configura una respuesta ok con el payload correspondiente y envia la respuesta al cliente
func SendResponse(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(Response{
		Code:    "0",
		Message: "",
		Payload: payload,
	})
	if err != nil {
		log.Printf("No fue posible generar JSON para enviar respuesta. Detalle: %s", err.Error())
		SendResponseErr(w, http.StatusInternalServerError, Response{
			Code:    "-1",
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// SendResponseErr configura una respuesta nok y envia la respuesta al cliente
func SendResponseErr(w http.ResponseWriter, httpStatus int, resp Response) {
	if resp.Code == "0" {
		resp.Code = "-1"
	}

	json, _ := json.Marshal(resp)

	w.WriteHeader(httpStatus)
	w.Write(json)
}

// SendResponseBadReq construye la respuesta con codigo HTTP 400
func SendResponseBadReq(w http.ResponseWriter, resp Response) {
	SendResponseErr(w, http.StatusBadRequest, resp)
}

// SendResponseInternalErr construye la respuesta con codigo HTTP 500
func SendResponseInternalErr(w http.ResponseWriter, resp Response) {
	SendResponseErr(w, http.StatusInternalServerError, resp)
}
