package network

type NetAddr string

type Message struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan Message
	Connect(transport Transport) error
	SendMessage(to NetAddr, payload []byte) error
	Addr() NetAddr
}
