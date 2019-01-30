package framework

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type (
	//
	// DbConfig is a type that holds JSON-loaded
	// fields that are related to the connection
	// to the database; the credentials
	//
	DbConfig struct {
		DbFile string `json:"database_file"`
	}

	//
	// Database is a wrapper type around sql.DB
	//
	Database struct {
		Handle *sql.DB
	}
)

//
// NewDatabase loads credential information from config/database.json,
// and opens a database connection through TCP
//
func NewDatabase(initDatabase bool) *Database {
	var config DbConfig
	if err := LoadFromJSON("config/database.json", &config); err != nil {
		fmt.Println("NewDatabase() loading the config file failed:", err)
		return nil
	}
	database := new(Database)

	db, err := sql.Open("sqlite3", config.DbFile)
	if err != nil {
		fmt.Println("NewDatabase() sql.Open() failed:", err)
		return nil
	}
	if err = db.Ping(); err != nil {
		fmt.Println("NewDatabase() db.Ping() failed:", err)
		return nil
	}

	if initDatabase {
		log.Println("initDatabase = True, creating tables...")
		stmt := `CREATE TABLE user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			discord_id TEXT NOT NULL,
			hunter_name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS user_guilds (
			user_id INTEGER NOT NULL,
			guild_id TEXT NOT NULL,
			PRIMARY KEY (user_id, guild_id),
			FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
		);`

		if _, err := db.Exec(stmt); err != nil {
			log.Fatalf("%q: %s\n", err, stmt)
			return nil
		}
		log.Println("tables created")
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
