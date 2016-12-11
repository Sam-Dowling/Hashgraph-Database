package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Sender    Peer
	code      int
	data      string
	PeerCount int
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
