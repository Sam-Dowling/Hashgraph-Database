package main

import (
	"crypto/ecdsa"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

type Peer struct {
	ID        int
	IP        string
	Port      int
	PublicKey ecdsa.PublicKey
}

type Message struct {
	Code int64
	Data interface{}
}

type EventCount struct {
	Count []int
}

type Events struct {
	Events []Event
}

var Network = []Peer{}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func AddPeer(p Peer) {
	Network = append(Network, p)
}

func GetRandomPeer() Peer {
	length := len(Network)
	if length > 0 {
		return Network[rand.Intn(length)]
	}
	return Peer{}
}

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
			gob.Register()
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
