package framework

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//
// DbConfig is a type that holds JSON-loaded
// fields that are related to the connection
// to the database; the credentials
//
type DbConfig struct {
	// Database fields
	DbHost string `json:"database_host"`
	DbPort string `json:"database_port"`
	DbUser string `json:"database_user"`
	DbName string `json:"database_name"`
	DbPass string `json:"database_pass"`
}

//
// Database is a wrapper type around sql.DB
//
type Database struct {
	Handle *sql.DB
}

//
// NewDatabase loads credential information from config/database.json,
// and opens a database connection through TCP
//
func NewDatabase() *Database {
	var config DbConfig
	if err := LoadFromJSON("config/database.json", &config); err != nil {
		fmt.Println("NewDatabase() loading the config file failed:", err)
		return nil
	}
	database := new(Database)

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName))

	if err != nil {
		fmt.Println("NewDatabase() sql.Open() failed:", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		fmt.Println("NewDatabase() db.Ping() failed:", err)
		return nil
	}
	database.Handle = db
	return database
}

//
// Close acts as a wrapper method for
// sql.Close()
//
func (db *Database) Close() (err error) {
	if db.Handle == nil {
		return
	}
	err = db.Handle.Close()
	return
}
