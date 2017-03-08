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
		message, ok := data.Data.(EventCount)
		if ok {
			//Debug
			fmt.Println(message.Transaction)
		}

	case 1:
		message, ok := data.Data.(Events)
		if ok {
			//Debug
			fmt.Println(message.Transaction)
		}
	}
	conn.Close()
}

func sendMessage(msg Message, p Peer) {
	conn, err := net.Dial("tcp", p.String())
	if err == nil {
		encoder := gob.NewEncoder(conn)
		switch msg.Code {
		case 0: //Gossip event counts
			gob.Register(EventCount{})
		case 1: //Events
			gob.Register(Events{})
		}
		encoder.Encode(&msg)
		conn.Close()
	}
}

func Gossip() {
	p := GetRandomPeer()
	if p != (Peer{}) {
		sendMessage(Message{0, getEventCounts()}, p)
	}
}

func sendEvents(int count, p Peer) {
	sendMessage(Message{1, Events{}}, p)
}
