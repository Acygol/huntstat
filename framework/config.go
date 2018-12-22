package framework

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token 		string 	`json:"token"`
	Prefix		string 	`json:"prefix"`

	// Database fields
	DbHost		string	`json:"database_host"`
	DbPort		string	`json:"database_port"`
	DbUser		string	`json:"database_user"`
	DbName		string 	`json:"database_name"`
	DbPass 		string 	`json:"database_pass"`

	DbHandle	*sql.DB
}

func Init(fileName string) *Config {
	conf := LoadConfig(fileName)

	// prepare database connection
	dbHandle, err := OpenDatabase(conf)
	if err != nil {
		fmt.Println("error preparing database connection,", err)
		return nil
	}
	conf.DbHandle = dbHandle
	fmt.Println("DbHandle in Init(): ", conf.DbHandle)
	return conf
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
