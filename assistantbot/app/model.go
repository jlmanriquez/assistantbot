package app

// HumanResourceTalk almacena los datos de una conversacion enviada por el usuario
type HumanResourceTalk struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}
