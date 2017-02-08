package main

type Message struct {
	Code int64
	Data interface{}
}

type PeerData struct {
	Peers []Peer
}

type TransactionData struct {
	Transaction string
}
