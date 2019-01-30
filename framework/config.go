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
func NewConfig(initDatabase bool) *Config {
	var conf Config
	if err := LoadFromJSON("config/config.json", &conf); err != nil {
		log.Fatal("error loading config file,", err)
		return nil
	}
	conf.Database = *NewDatabase(initDatabase)
	return &conf
}
