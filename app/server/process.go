package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/simon/mengine/app/handler"
	"github.com/simon/mengine/app/service"
	"github.com/simon/mengine/infrastructure/engine"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type process struct {
	gin *gin.Engine
}

var WireHttpProcessSet = wire.NewSet(NewProcess)

func NewProcess() engine.Process {
	return &process{
		gin: gin.Default(),
	}
}

func (p *process) Name() string {
	return `http`
}

func (p *process) Global() bool {
	return false
}

func (p *process) Prepare(engine engine.Engine) error {
	service.DefaultAi = service.NewAi("")
	handler.DefaultHandler = &handler.Handler{Ai: service.DefaultAi}

	registerRoutes(p.gin, handler.DefaultHandler)
	return nil
}

func (p *process) Shutdown(engine engine.Engine) error {
	return nil
}

func (p *process) Cobra() *cobra.Command {
	return &cobra.Command{
		Use:   `http.serve`,
		Short: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetOutput(os.Stdout)

			return p.gin.Run(":28080")
		},
	}
}
