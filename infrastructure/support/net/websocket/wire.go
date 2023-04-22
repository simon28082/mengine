package websocket

import (
	"github.com/google/wire"
	"github.com/gorilla/websocket"
)

func ProvideGorilla(id string, conn *websocket.Conn) *socket {
	panic(wire.Build(NewWebSocket))
}
