package main

import (
	"fmt"
	"math/rand"
)

type Peer struct {
	IP   string
	Port int
}

var Peers = []Peer{}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func AddPeer(p Peer) {
	if !IsKnownPeer(p) {
		Peers = append(Peers, p)
	}
}

func GetRandomPeer() Peer {
	return Peers[rand.Intn(len(Peers))]
}

func GetPeerCount() int {
	return len(Peers)
}

func IsKnownPeer(p Peer) bool {
	for _, peer := range Peers {
		if p == peer {
			return true
		}
	}
	return false
}
