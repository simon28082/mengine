package cmd

import "github.com/google/wire"

type provider struct {
}

func (p provider) Name() string {
	return `cmd`
}

func (p provider) Register() {
	wire.Build(NewEvent, NewGreeter, NewMessage)
}

func (p provider) Boot() {
}
