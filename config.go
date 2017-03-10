package main

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	ID         int
	IP         string
	Port       int
	PrivateKey ecdsa.PrivateKey
	Network    []Peer
}

var configFile = "config.toml"

var GlobalConfig = config{}

func VerifyConfigFile(filename string) {
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", filename)
	} else {
		configFile = filename
	}
}

// Reads info from config file
func ReadConfig() {

	if len(os.Args) > 1 {
		VerifyConfigFile(os.Args[1])
	} else {
		VerifyConfigFile(configFile)
	}
	if _, err := toml.DecodeFile(configFile, &GlobalConfig); err != nil {
		log.Fatal(err)
	}
	for _, peer := range GlobalConfig.Network {
		AddPeer(peer)
	}
}

func SaveConfig() {
	GlobalConfig.Network = Network

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}
	encoder := toml.NewEncoder(file)
	encoder.Encode(&GlobalConfig)
	file.Close()
}
