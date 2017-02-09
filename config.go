package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	IP    string
	Port  int
	Peers []Peer
}

var ConfigFile = "config.toml"

var GlobalConfig = Config{}

func VerifyConfigFile(filename string) {
	_, err := os.Stat(ConfigFile)
	if err != nil {
		log.Fatal("Config file is missing: ", filename)
	} else {
		ConfigFile = filename
	}
}

// Reads info from config file
func ReadConfig() {

	if len(os.Args) > 1 {
		VerifyConfigFile(os.Args[1])
	} else {
		VerifyConfigFile(ConfigFile)
	}
	if _, err := toml.DecodeFile(ConfigFile, &GlobalConfig); err != nil {
		log.Fatal(err)
	}

	for _, peer := range GlobalConfig.Peers {
		AddPeer(peer)
	}
}

func SaveConfig() {
	GlobalConfig.Peers = PeerList

	file, err := os.Create(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	encoder := toml.NewEncoder(file)
	encoder.Encode(&GlobalConfig)
	file.Close()
}
