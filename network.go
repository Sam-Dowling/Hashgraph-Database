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
	data := &Message{}
	dec.Decode(data)

	switch data.Code {
	case 0:
		peerList := data.Data.(PeerData)
		for _, peer := range peerList.Peers {
			AddPeer(peer)
		}
	case 1:
		message := data.Data.(TransactionData)
		fmt.Println(message.Transaction)
	}

	conn.Close()
}

func sendMessage(msg Message, p Peer) {
	conn, err := net.Dial("tcp", p.String())
	if err == nil {
		encoder := gob.NewEncoder(conn)
		switch msg.Code {
		case 0:
			gob.Register(PeerData{})
		case 1:
			gob.Register(TransactionData{})
		}
		encoder.Encode(&msg)
		conn.Close()
	}
}

func PeerExchange() {
	sendMessage(Message{0, PeerData{PeerList}}, GetRandomPeer())
}

func Gossip() {
	sendMessage(Message{1, TransactionData{strconv.Itoa(GlobalConfig.Port)}}, GetRandomPeer())
}
