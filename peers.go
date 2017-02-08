package main

import (
	"fmt"
	"math/rand"
)

type Peer struct {
	IP   string
	Port int
}

var PeerList = []Peer{}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func AddPeer(p Peer) {
	if !IsKnownPeer(p) {
		PeerList = append(PeerList, p)
	}
}

func GetRandomPeer() Peer {
	return PeerList[rand.Intn(len(PeerList))]
}

func GetPeerCount() int {
	return len(PeerList)
}

func IsKnownPeer(p Peer) bool {
	for _, peer := range PeerList {
		if p == peer {
			return true
		}
	}
	return false
}
