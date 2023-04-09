package server

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/simon/mengine/app/handler"
	"net/http"
	"time"
)

func registerRoutes(app *gin.Engine, handler *handler.Handler) {
	app.Use(cors.Default())
	spew.Dump("######")
	app.POST(`/chat`, handler.Chat)
	app.GET(`/chat-stream`, handler.ChatStream)
	app.GET(`/test`, func(ctx *gin.Context) {
		ctx.String(200, "hello")
		ctx.Next()
	})

	app.GET("/stream-html", func(ctx *gin.Context) {
		ctx.Header("content-type", "text/html; charset=utf-8")
		ctx.HTML(200, "index.html", map[string]interface{}{})
	})
	app.GET("/stream", func(ctx *gin.Context) {
		w := ctx.Writer
		w.Header().Set("Content-Type", "text/event-stream")

		// Write the initial response to the client
		fmt.Fprintf(w, "data: Connected\n\n")

		// Loop forever and send a message to the client every second
		for {
			// Generate a random message
			message := fmt.Sprintf("data: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

			// Write the message to the client
			fmt.Fprintf(w, message)

			// Flush the response writer to ensure that the message is sent immediately
			w.(http.Flusher).Flush()

			// Wait for one second before sending the next message
			time.Sleep(1 * time.Second)
		}
	})
}

func e() {

}
