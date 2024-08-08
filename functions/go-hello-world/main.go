package handler

import (
	"context"
	"os"

	"github.com/open-runtimes/types-for-go/v4"
	openai "github.com/sashabaranov/go-openai"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

func Main(Context *types.Context) types.ResponseOutput {

	openAiKey := os.Getenv("OPENAI_KEY")

	openAiClient := openai.NewClient(openAiKey)

	if Context.Req.Method == "GET" {
		return Context.Res.Text("Hello, World!", 200, nil)
	}

	if Context.Req.Method == "POST" {
		requestBody := Context.Req.BodyJson()

		prompt := requestBody["prompt"].(string)

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
			}, 500, nil)
		}

		return Context.Res.Json(map[string]interface{}{
			"ok":       true,
			"response": completion.Choices[0].Message.Content,
		}, 200, nil)
	}

	// `Res.Json()` is a handy helper for sending JSON
	return Context.Res.Json(map[string]interface{}{
		"motto":       "Build like a team of hundreds_",
		"learn":       "https://appwrite.io/docs",
		"connect":     "https://appwrite.io/discord",
		"getInspired": "https://builtwith.appwrite.io",
	}, 200, nil)
}
