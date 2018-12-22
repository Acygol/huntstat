package framework

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token 	string `json:"token"`
	Prefix	string `json:"prefix"`
}

//
// LoadConfig reads from config.json to populate
// the Bot struct
func LoadConfig(fileName string) *Config {
	// Open config file
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("failed to open config file, ", err)
		return nil
	}
	// Populate configIns fields and return its address
	var config Config
	json.Unmarshal(body, &config)
	return &config
}

