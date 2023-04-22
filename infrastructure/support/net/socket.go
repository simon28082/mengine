package net

import "net"

type Socket interface {
	net.Conn

	Id() string

	Close() error
}
