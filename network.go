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
*/

type Peer struct {
	IP        string
	Port      int
	PublicKey rsa.PublicKey
}

type Message struct {
	Address Peer
	Code    int
	Data    interface{}
}

type EventCount struct {
	Count map[string]int
}

type Events struct {
	Head      string
	EventList map[string]Event
}

var Network = make(map[string]Peer)

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

	AddPeer(data.Address)

	switch data.Code {
	case 0: // 0 : EventCount - Gossip EventCount
		message, ok := data.Data.(EventCount)
		if ok {
			request := CalcEventsToRequest(message.Count)
			sendMessage(Message{Self, 1, EventCount{request}}, data.Address)
		}
		break

	case 1: // 1 : EventCount - request Events EventCount
		message, ok := data.Data.(EventCount)
		if ok {
			e := CollectEventsToSend(message.Count)

			sendMessage(Message{Self, 2, Events{Head, e}}, data.Address)
		}
		break

	case 2: // 2 : Events - last n Events from each node defined in 1:EventCount
		message, ok := data.Data.(Events)
		fmt.Println("Received Events")
		if ok {
			ParseEvents(message)
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
			break
		case 1: //request event counts
			gob.Register(EventCount{})
			break
		case 2: //Events
			gob.Register(Events{})
			break
		}
		encoder.Encode(&msg)
		conn.Close()
	}
}

func Gossip() {
	p := GetRandomPeer()
	if p != (Peer{}) {
		fmt.Println("Gossiping")
		sendMessage(Message{Self, 0, EventCount{GetEventCounts()}}, p)
	}
}
