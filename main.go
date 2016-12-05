package main

import "os"

func main() {

	if len(os.Args) > 1 {
		GloblConfig = ReadConfig(os.Args[1])
	}

	//go StartListening()

	// for {
	// 	time.Sleep(time.Second * 2)
	//
	// 	Gossip()
	// }

}
