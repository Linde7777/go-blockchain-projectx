package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	peers     map[NetAddr]*LocalTransport
	lock      sync.RWMutex
	consumeCh chan Message
}

var _ Transport = &LocalTransport{}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		peers:     make(map[NetAddr]*LocalTransport),
		lock:      sync.RWMutex{},
		consumeCh: make(chan Message, 1024),
	}
}

// Consume get the message from the transport
//
// example usage:
// recvMessage := <-tran.Consume()
// print(recvMessage.Payload)
// print(recvMessage.From)
func (l *LocalTransport) Consume() <-chan Message {
	return l.consumeCh
}

func (l *LocalTransport) Connect(t Transport) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.peers[t.Addr()] = t.(*LocalTransport)

	return nil
}

func (l *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	l.lock.RLock()
	defer l.lock.RUnlock()

	transport, ok := l.peers[to]
	if !ok {
		return fmt.Errorf("fail to get transport, related NetAddr:%s", to)
	}
	transport.consumeCh <- Message{
		From:    l.addr,
		Payload: payload,
	}

	return nil
}

func (l *LocalTransport) Addr() NetAddr {
	return l.addr
}
