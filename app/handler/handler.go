package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simon/mengine/app/service"
)

var DefaultHandler *Handler

type Handler struct {
	Ai *service.Ai
}

func (h *Handler) Chat(ctx *gin.Context) {
	fmt.Println(ctx.Query(`message`))
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
