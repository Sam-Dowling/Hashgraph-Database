package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

/*
-API-
code : struct - description

0 : EventCount - Gossip EventCount

1 : EventCount - request Events EventCount
	<- 2 : Events - last n Events from each node defined in 1:EventCount

3 : Peer - handshake
*/

type Peer struct {
	IP        string
	Port      int
	PublicKey rsa.PublicKey
}

type Message struct {
	Code int
	Data interface{}
}

type EventCount struct {
	Count map[string]int
}

type Events struct {
	Events map[string]Event
}

var Network = map[string]Peer{}

func (p Peer) toString() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func AddPeer(p Peer) {
	Network[p.toString()] = p
}

func GetRandomPeer() Peer {
	i := rand.Intn(len(Network))

	for _, v := range Network {
		if i == 0 {
			return v
		}
		i--
	}
	return Peer{} // empty Network
}

func StartListening() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(GlobalConfig.Port))
	if err == nil {
		fmt.Println("Started Listening on " + Self.toString())
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
	case 0: // 0 : EventCount - Gossip EventCount
		message, ok := data.Data.(EventCount)
		if ok {
			fmt.Println(message.Count)
		}
		break

	case 1: // 1 : EventCount - request Events EventCount
		message, ok := data.Data.(EventCount)
		if ok {
			//Debug
			fmt.Println(message.Count)
		}
		break

	case 2: // 2 : Events - last n Events from each node defined in 1:EventCount
		message, ok := data.Data.(Events)
		if ok {
			//Debug
			fmt.Println(message.Events)
		}
		break
	}
	conn.Close()
}

func sendMessage(msg Message, p Peer) {
	conn, err := net.Dial("tcp", p.toString())
	if err == nil {
		encoder := gob.NewEncoder(conn)
		switch msg.Code {
		case 0: //Gossip event counts
			gob.Register(EventCount{})
		case 1: //request event counts
			gob.Register(EventCount{})
		case 2: //Events
			gob.Register(Events{})
		}
		encoder.Encode(&msg)
		conn.Close()
	}
}

func Gossip() {
	p := GetRandomPeer()
	if p != (Peer{}) {
		sendMessage(Message{0, EventCount{}}, p)
	}
}
