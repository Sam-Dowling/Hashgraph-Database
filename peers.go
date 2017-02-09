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
	length := len(PeerList)
	if length > 0 {
		return PeerList[rand.Intn(length)]
	}
	return Peer{}
}

func IsKnownPeer(p Peer) bool {
	if p == (Peer{GlobalConfig.IP, GlobalConfig.Port}) {
		return true
	}
	for _, peer := range PeerList {
		if p == peer {
			return true
		}
	}
	return false
}
