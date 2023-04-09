package server

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/simon/mengine/app/handler"
)

func registerRoutes(app *gin.Engine, handler *handler.Handler) {
	spew.Dump("######")
	app.POST(`/chat`, handler.Chat)
	app.POST(`/chat-stream`, handler.Chat)
	app.GET(`/test`, func(ctx *gin.Context) {
		ctx.String(200, "hello")
		ctx.Next()
	})
}
