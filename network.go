package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
)

func StartListening() {
	ln, _ := net.Listen("tcp", string(GlobalConfig.Port))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	address := conn.RemoteAddr().String()
	fmt.Print(address, "> ", string(message))
	conn.Close()
}

func sendMessage(msg Message, p Peer) {
	conn, err := net.Dial("tcp", p.String())
	if err == nil {
		enc := gob.NewEncoder(conn)
		enc.Encode(msg)
		conn.Close()
	}
}

func Gossip() {
	sendMessage("Gossip", GetRandomPeer())
}
