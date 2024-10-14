package main

import (
	"github.com/Linde7777/go-blockchain-projectx/network"
	"time"
)

func main() {
	tran1 := network.NewLocalTransport("local1")
	tran2 := network.NewLocalTransport("local2")
	err := tran1.Connect(tran2)
	if err != nil {
		panic(err)
	}
	err = tran2.Connect(tran1)
	if err != nil {
		panic(err)
	}
	opts1 := network.ServerOpts{Transports: []network.Transport{tran1}}
	server1 := network.NewServer(opts1)

	go func() {
		for {
			err := tran2.SendMessage(tran1.Addr(), []byte("blah"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	server1.Start()

}
