package main

import (
	"os"
	"time"
)

func main() {

	if len(os.Args) > 1 {
		GlobalConfig = ReadConfig(os.Args[1])
	}

	go StartListening()

	for {
		time.Sleep(time.Second * 5)

		Gossip()
	}

}
