package main

import (
	"encoding/json"
	"fmt"
)

/*

Request Peer Data : 102
Peer Data         : 103


*/

type Message struct {
	Sender    Peer
	Code      int
	data      string
	PeerCount int
}

func createGossip(msg string) string {
	return
}

func (m Message) MessagetoJson() string {
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		return string(json)
	}
}

func (j string) JsontoMessage() Message {
	s, err := json.Unmarshal(j, &Message)
	if err != nil {
		fmt.Println(err)
		return Message{}
	} else {
		return s
	}
}
