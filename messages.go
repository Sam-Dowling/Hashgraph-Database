package main

type Message struct {
	Code int64
	Data string
}

type PeerData struct {
	Peers []Peer
}
