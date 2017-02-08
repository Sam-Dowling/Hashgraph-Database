package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) > 1 {
		GlobalConfig = ReadConfig(os.Args[1])
	}

	fmt.Println("STARTING: ", GlobalConfig.Port)
	fmt.Println(PeerList)

	go StartListening()

	for i := 0; i < 50; i++ {
		time.Sleep(time.Second * 1)

		Gossip()
		PeerExchange()
	}
	fmt.Println(PeerList)

}
