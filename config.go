package main

import (
	"crypto/rsa"
	"crypto/x509"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var PrivateKey *rsa.PrivateKey

type config struct {
	IP              string
	Port            int
	PrivateKeyBytes []byte
	Network         map[string]Peer
}

var configFile = "config.toml"

var GlobalConfig = config{}

var Self = Peer{}

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

	//AddPeer(Self)

	priv, err := x509.ParsePKCS1PrivateKey(GlobalConfig.PrivateKeyBytes)
	if err != nil {
		//no key present? generate one
		priv = GenKey()
	}
	PrivateKey = priv
	Self = Peer{GlobalConfig.IP, GlobalConfig.Port, PrivateKey.PublicKey}
}

func SaveConfig() {
	GlobalConfig.Network = Network

	file, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
	}

	GlobalConfig.PrivateKeyBytes = x509.MarshalPKCS1PrivateKey(PrivateKey)

	encoder := toml.NewEncoder(file)
	encoder.Encode(&GlobalConfig)
	file.Close()
}
