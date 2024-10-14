package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalTransport_Addr(t *testing.T) {
	t1 := NewLocalTransport("t1")
	assert.Equal(t, t1.Addr(), NetAddr("t1"))
}

func TestLocalTransport_Connect(t *testing.T) {
	t1 := NewLocalTransport("t1")
	t2 := NewLocalTransport("t2")
	assert.Nil(t, t1.Connect(t2))
	assert.Nil(t, t2.Connect(t1))

	assert.Equal(t, t1.peers[t2.Addr()], t2)
	assert.Equal(t, t2.peers[t1.Addr()], t1)
}

func TestLocalTransport_Consume(t *testing.T) {
	t1 := NewLocalTransport("t1")
	t2 := NewLocalTransport("t2")
	assert.Nil(t, t1.Connect(t2))
	assert.Nil(t, t2.Connect(t1))

	sentMessage := []byte("red dead redemption")
	assert.Nil(t, t1.SendMessage(t2.Addr(), sentMessage))

	recvMessage := <-t2.Consume()
	assert.Equal(t, recvMessage.Payload, sentMessage)
	assert.Equal(t, recvMessage.From, t1.Addr())
}
