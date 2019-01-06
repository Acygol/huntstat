package framework

import (
	"log"
)

//
// Config is a type that holds JSON-loaded
// information related to the bot, such as
// the Discord connection token and command
// prefix
//
type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`

	Database Database
}

//
// NewConfig creates a new instance of Config
// initiating its fields along with a database
// connection. It then returns a pointer
// to the newly created config
//
func NewConfig() *Config {
	conf := loadConfig()
	conf.Database = *NewDatabase()
	return conf
}

//
// loadConfig reads from config.json to populate
// the Bot struct
//
func loadConfig() *Config {
	var config Config
	if err := LoadFromJSON("config/config.json", &config); err != nil {
		log.Fatal("error loading config file,", err)
		return nil
	}
	return &config
}
