package handler

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/simon/mengine/app/service"
	"io"
)

var DefaultHandler *Handler

type Handler struct {
	Ai *service.Ai
}

func (h *Handler) ChatStream(ctx *gin.Context) {
	//message, ok := ctx.GetPostForm(`message`)
	//if !ok {
	//	ctx.Error(errors.NewDefault(`message not found`))
	//	return
	//}
	message := "your is chatgpt?"

	//ctx.Header("Content-Type", "text/event-stream")
	//ctx.Header("Cache-Control", "no-cache")
	//ctx.Header("Connection", "keep-alive")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	//for {
	//	// Write the message to the client
	//	spew.Dump("++++++++======++++++++")
	//	ctx.Writer.Write([]byte())
	//	// Flush the response writer to ensure that the message is sent immediately
	//	ctx.Writer.Flush()
	//	time.Sleep(time.Second)
	//}

	//for {
	//	select {
	//	case <-ctx.Writer.CloseNotify():
	//		// Stop streaming when the client disconnects
	//		return
	//	default:
	//		log.Println("++++++333333")
	//		// Write the data to the client
	//		ctx.Stream(func(w io.Writer) bool {
	//			_, err := w.Write([]byte(time.Now().String()))
	//			return err == nil
	//		})
	//	}
	//}

	//ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	resps, err := h.Ai.CompletionStream(ctx, openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	spew.Dump("++++++++++++++++", resps, len(resps))

	ctx.Stream(func(w io.Writer) bool {
		if resp, ok := <-resps; ok {
			ctx.SSEvent("message", resp.Choices[0].Delta.Content)
			return true
		}

		return false
	})

	//if resp.Choices[0].FinishReason == "stop" {
	//	return false
	//}

	//for resp := range resps {
	//	// Write the message to the client
	//	fmt.Fprintf(ctx.Writer, resp.Choices[0].Delta.Content)
	//	spew.Dump("++++++++======++++++++")
	//
	//	// Flush the response writer to ensure that the message is sent immediately
	//	ctx.Writer.(http.Flusher).Flush()
	//
	//}
}

func (h *Handler) Chat(ctx *gin.Context) {
	fmt.Println(ctx.GetPostForm(`message`))
	//resp, err := h.Ai.Completion(ctx, gogpt.CompletionRequest{
	//	Prompt: "再多说一点？",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.Error(err)
	//
	//	return
	//}
	//ctx.String(200, resp)
	//ctx.Next()
}
