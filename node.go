package main

import "time"

type Transaction struct {
	timestamp time.Time
	action    int // 0:create, 1:read, 2:update
	key       string
	value     string
}

type Event struct {
	timestamp    time.Time
	transactions []Transaction
	selfParent   string
	otherParent  string
}

type Node struct {
	ID        int
	Network   []Peer
	hashgraph map[string]Event
}

func (n *Node) Run() {
	//sig, event := n.createEvent()
	//n.addEvent(sig, event)

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 2)
		//n.sync()
	}

}

// func (n *Node) createEvent() {
//
// }
