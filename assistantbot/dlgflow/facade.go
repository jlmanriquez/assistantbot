package dlgflow

import (
	"context"
	"fmt"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

const (
	projectID    = "segurito-6acde"
	languageCode = "es"
)

// SendIntent envia un texto a Dialogflow para realizar la conversacion
func SendIntent(text, sessionID string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer sessionClient.Close()

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}
