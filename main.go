package main

import (
	"os"
	"time"
)

func main() {

	if len(os.Args) > 1 {
		GlobalConfig = ReadConfig(os.Args[1])
	}

	AddPeer(Peer{"127.0.0.1", 9000})

	go StartListening()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)

		Gossip()
	}

}
