package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	timestamp time.Time
	action    int // 0:create, 1:read, 2:update
	key       string
	value     string
}

type Event struct {
	creator      string
	timestamp    time.Time
	transactions []Transaction
	selfParent   string
	otherParent  string
}

var Hashgraph map[string]Event
var transactions []Transaction
var head string

func Run() {
	Hashgraph = make(map[string]Event)

	sig, e := createEvent(nil, "", "")
	addEvent(sig, e)

	StartListening()
	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 2)
		Gossip()
	}

	// createTransaction(0, "127.0.0.1", "example.com")
	// createTransaction(1, "123.456.789.0", "another.net")

}

func getEventCounts(hg map[string]Event) map[string]int {
	count := make(map[string]int)
	for _, v := range hg {
		count[v.creator]++
	}
	return count
}

func (e Event) toString() string {
	return fmt.Sprintf("%v", e)
}

func createTransaction(action int, key string, value string) {
	transactions = append(transactions, Transaction{time.Now(), action, key, value})
}

func createEvent(trans []Transaction, selfParent string, otherParent string) (string, Event) { // (events signature, event)
	e := Event{Self.toString(), time.Now(), trans, selfParent, otherParent}
	eventStr := e.toString()
	return CalcSignature(eventStr), e
}

func addEvent(sig string, e Event) {
	if e.creator == Self.toString() {
		head = sig
	}
	Hashgraph[sig] = e
}
