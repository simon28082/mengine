package cmd

import (
	"github.com/simon/mengine/infrastructure/cmd"
	"github.com/simon/mengine/infrastructure/support/chain"
	"github.com/simon/mengine/server/http"
)

func main() {
	var (
		cmdFunc = func(next chain.ChainFunc) chain.ChainFunc {
			cmd.NewProvider()
			return next
		}
		httpFunc = func(next chain.ChainFunc) chain.ChainFunc {
			http.NewProvider()
			return next
		}
	)

	chain.BuildChain(cmdFunc, httpFunc)
}
