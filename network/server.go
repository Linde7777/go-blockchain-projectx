package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	opts       ServerOpts
	messageCh  chan Message
	shutdownCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		opts:       opts,
		messageCh:  make(chan Message),
		shutdownCh: make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

quitForLoop:
	for {
		select {
		case message := <-s.messageCh:
			fmt.Printf("receive message: %+v\n", message)
		case <-s.shutdownCh:
			break quitForLoop
		case <-ticker.C:
			fmt.Println("do something every 5 second")
		}
	}
}

func (s *Server) initTransports() {
	for _, tran := range s.opts.Transports {
		go func(tran Transport) {
			for message := range tran.Consume() {
				s.messageCh <- message
			}
		}(tran)
	}
}
