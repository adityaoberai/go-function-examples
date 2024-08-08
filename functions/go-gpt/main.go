package handler

import (
	"context"
	"os"

	"github.com/open-runtimes/types-for-go/v4/openruntimes"
	openai "github.com/sashabaranov/go-openai"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

// This Appwrite function will be executed every time your function is triggered
func Main(Context openruntimes.Context) openruntimes.Response {
	openAiKey := os.Getenv("OPENAI_KEY")

	openAiClient := openai.NewClient(openAiKey)

	if Context.Req.Method == "GET" {
		return Context.Res.Text("Hello, World!", Context.Res.WithStatusCode(200), nil)
	}

	if Context.Req.Method == "POST" {
		var requestBody RequestBody
		err := Context.Req.BodyJson(&requestBody)

		if err != nil {
			Context.Error(err)
			Context.Res.WithStatusCode(400)
			return Context.Res.Json(map[string]interface{}{
				"ok":    false,
				"error": "Missing request body",
			}, Context.Res.WithStatusCode(400), nil)
		}

		prompt := requestBody.Prompt

		completion, err := openAiClient.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
			},
		)

		if err != nil {
			Context.Error(err)
			return Context.Res.Json(map[string]interface{}{
				"ok":    false,
				"error": err,
			}, Context.Res.WithStatusCode(500), nil)
		}

		return Context.Res.Json(map[string]interface{}{
			"ok":       true,
			"response": completion.Choices[0].Message.Content,
		}, Context.Res.WithStatusCode(200), nil)
	}

	return Context.Res.Json(map[string]interface{}{
		"ok":    false,
		"error": "Bad request",
	}, Context.Res.WithStatusCode(400), nil)
}
