package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type Configuration struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func LoadConfig() Configuration {
	conf := Configuration{}
	configFile, err := os.Open("./config/config.json")

	defer configFile.Close()

	if err != nil {
		fmt.Println("Error loading config.json",err.Error())
		panic("Error loading config.json")
	}

	err = json.NewDecoder(configFile).Decode(&conf)

	if err != nil {
		panic("Error loading config object from config.json")
	}

	return conf
}
