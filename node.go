package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transaction struct {
	Timestamp time.Time
	Action    int // 0:create, 1:read, 2:update
	Key       string
	Value     string
}

type Event struct {
	Creator      string
	Timestamp    time.Time
	Transactions []Transaction
	SelfParent   string
	OtherParent  string
}

var order = make(map[string][]*string)

var Hashgraph = make(map[string]Event)
var transactions []Transaction
var Head string

func Run() {
	//createTransaction(0, "127.0.0.1", "example.com")
	sig, e := createEvent("0", "0")
	addEvent(sig, e)

	go StartListening()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		Gossip()
		fmt.Println(len(Hashgraph))
	}

	b, err := json.MarshalIndent(Hashgraph, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

}

func ParseEvents(event Events) {
	for k, v := range event.EventList {
		if VerifySignature(Network[v.Creator].PublicKey, k, v.toString()) { // Valid signature of event
			addEvent(k, v)
		}
	}
	sig, e := createEvent(Head, event.Head)
	addEvent(sig, e)
}

func CollectEventsToSend(eventCounts map[string]int) map[string]Event {
	e := make(map[string]Event)
	for k, v := range eventCounts {
		j := len(order[k]) - v
		for i := j; i < len(order[k]); i++ {
			e[*order[k][i]] = Hashgraph[*order[k][i]]
		}
	}
	return e
}

func GetEventCounts() map[string]int {
	count := make(map[string]int)
	for _, v := range Hashgraph {
		count[v.Creator]++
	}
	return count
}

func CalcEventsToRequest(theirCount map[string]int) map[string]int {
	request := make(map[string]int)
	ourCount := GetEventCounts()

	for k, v := range theirCount {
		if ourCount[k] < v {
			request[k] = v - ourCount[k]
		}
	}
	return request
}

func (e Event) toString() string {
	return fmt.Sprintf("%v", e)
}

func createTransaction(action int, key string, value string) {
	transactions = append(transactions, Transaction{time.Now(), action, key, value})
}

func createEvent(selfParent string, otherParent string) (string, Event) { // (events signature, event)
	e := Event{Self.toString(), time.Now(), transactions, selfParent, otherParent}
	transactions = nil //clear transactions
	eventStr := e.toString()
	return CalcSignature(eventStr), e
}

func addEvent(sig string, e Event) {
	if e.Creator == Self.toString() {
		Head = sig
	}
	order[e.Creator] = append(order[e.Creator], &sig)
	Hashgraph[sig] = e
}
