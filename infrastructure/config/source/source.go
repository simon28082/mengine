package source

import "io"

type Source interface {
	Read() (io.Reader, error)

	Close() error
}

type Notifier interface {
	Notify() (<-chan []byte, error)
}
