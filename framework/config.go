package framework

import (
	"log"
)

/*
// Config is a type that holds JSON-loaded
// information tokens related to the bot,
// such as the Discord connection token and
// the command prefix
*/
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

/*
// loadConfig reads from config.json to populate
// the Bot struct
*/

func loadConfig() *Config {
	var config Config
	err := LoadFromJson("config/config.json", &config)
	if err != nil {
		log.Fatal("error loading config file,", err)
		return nil
	}
	return &config
}

/*
func loadConfig(fileName string) *Config {
	// Open config file
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("failed to open config file,", err)
		return nil
	}
	// Populate config fields and return its address
	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println("failed to unmarshal JSON,", err)
		return nil
	}
	return &config
}
*/
