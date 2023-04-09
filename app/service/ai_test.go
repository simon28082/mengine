package service

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"testing"
)

func TestAi_Completion(t *testing.T) {
	ai := NewAi()
	ai.Completion(context.Background(), openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Your is chatgpt3.5?",
			},
		},
	})
}
