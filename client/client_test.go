package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/laszukdawid/anthropic-go-sdk/client"
	"github.com/stretchr/testify/assert"
)

var anthropicVersion string = "2023-06-01"
var API_KEY string = os.Getenv("API_KEY")

func TestAuthenticatedClient(t *testing.T) {
	fmt.Printf("API_KEY: %s\n", API_KEY)
	const modelId = client.Claude3Haiku20240307
	// Set up authenticated client
	var content client.InputMessage_Content
	content.FromInputMessageContent0("Hello, World!")

	var model client.Model
	model.FromModel1(modelId)

	body := client.CreateMessageParams{
		MaxTokens: 100,
		Messages: []client.InputMessage{
			{Role: "user", Content: content},
		},
		Model: model,
	}

	// Send POST /v1/messages request
	apiClient, err := client.NewClientWithResponses("https://api.anthropic.com")
	if err != nil {
		t.Fatal(err)
	}

	params := client.MessagesPostParams{AnthropicVersion: &anthropicVersion, XApiKey: &API_KEY}
	response, err := apiClient.MessagesPost(context.Background(), &params, body)
	if err != nil {
		t.Fatal(err)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Response Body: %s\n", string(bodyBytes))

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "200 OK", response.Status)

	var msg client.Message
	json.Unmarshal(bodyBytes, &msg)
	assert.Equal(t, client.MessageTypeMessage, msg.Type)
	assert.Equal(t, client.MessageRoleAssistant, msg.Role)

	outModelId, err := msg.Model.AsModel1()
	assert.Nil(t, err)
	assert.Equal(t, modelId, outModelId)

	outStopReason, err := msg.StopReason.AsMessageStopReason0()
	assert.Nil(t, err)
	assert.Equal(t, client.MessageStopReason0EndTurn, outStopReason)

}
