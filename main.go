package main

import (
	"fmt"
	"time"
)

func main() {

	ReadConfig()

	fmt.Println("STARTING: ", GlobalConfig.Port)
	fmt.Println(PeerList)

	go StartListening()

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 2)

		Gossip()
		PeerExchange()
	}
	fmt.Println(PeerList)
	SaveConfig()

}
