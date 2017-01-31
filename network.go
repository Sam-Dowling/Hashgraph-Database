package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

func StartListening() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(GlobalConfig.Port))
	if err == nil {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go handleConn(conn)
		}
	}
	fmt.Println("Shutting Down")

}

func handleConn(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	p := &Message{}
	dec.Decode(p)

	address := conn.RemoteAddr().String()
	fmt.Println(address, " > ", p)
	conn.Close()
}

func sendMessage(msg Message, p Peer) {
	conn, err := net.Dial("tcp", p.String())
	if err == nil {
		encoder := gob.NewEncoder(conn)
		encoder.Encode(&msg)
		conn.Close()
	}
}

func Gossip() {
	sendMessage(Message{5, "Hello"}, GetRandomPeer())
}
