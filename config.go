package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	IP   string
	Port int
}

var GlobalConfig = ReadConfig("config.toml")

// Reads info from config file
func ReadConfig(configfile string) Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
