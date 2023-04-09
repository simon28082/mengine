package service

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAi_Completion(t *testing.T) {
	ai := NewAi()
	resp, err := ai.Completion(context.Background(), openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "can you help me?",
			},
		},
	})
	assert.Nil(t, err)
	spew.Dump(resp)
}

func TestAi_CompletionStream(t *testing.T) {
	ai := NewAi()
	resps, err := ai.CompletionStream(context.Background(), openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "can you help me? i need help",
			},
		},
	})
	assert.Nil(t, err)
	spew.Dump(resps)
	for r := range resps {
		spew.Dump(r)
	}
}
