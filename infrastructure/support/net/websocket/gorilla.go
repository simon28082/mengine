package websocket

import (
	"github.com/gorilla/websocket"
	"net"
	"sync"
	"time"
)

type socket struct {
	id           string
	writeLock    sync.RWMutex
	socketLock   sync.Mutex
	conn         *websocket.Conn
	writeMessage chan []byte
	wg           sync.WaitGroup
	readTimeout  time.Time
	writeTimeout time.Time
}

func NewWebSocket(id string, conn *websocket.Conn) *socket {
	return &socket{
		writeMessage: make(chan []byte, 1),
		conn:         conn,
		id:           id,
	}
}

func (s *socket) Id() string {
	return s.id
}

func (s *socket) Read(b []byte) (n int, err error) {
	if !s.readTimeout.IsZero() {
		s.conn.SetReadDeadline(s.readTimeout)
	}
	_, message, err := s.conn.ReadMessage()
	if err != nil {
		return 0, err
	}

	return copy(b, message), nil
}

func (s *socket) LocalAddr() net.Addr {
	return s.conn.LocalAddr()
}

func (s *socket) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *socket) SetDeadline(t time.Time) error {
	s.readTimeout = t
	s.writeTimeout = t
	return nil
}

func (s *socket) SetReadDeadline(t time.Time) error {
	s.readTimeout = t
	return nil
}

func (s *socket) SetWriteDeadline(t time.Time) error {
	s.writeTimeout = t
	return nil
}

func (s *socket) Write(message []byte) (n int, err error) {
	if !s.writeTimeout.IsZero() {
		s.conn.SetWriteDeadline(s.writeTimeout)
	}
	if err := s.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return 0, err
	}
	return len(message), nil
}

func (s *socket) Close() error {
	s.socketLock.Lock()
	defer s.socketLock.Unlock()

	if s.conn != nil {
		return s.conn.Close()
	}

	return nil
}
