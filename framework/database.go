package framework

import (
	"database/sql"
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
func NewDatabase() (Database, error) {
	var config DbConfig
	if err := LoadFromJSON("config/database.json", &config); err != nil {
		return Database{}, err
	}
	database := new(Database)

	db, err := sql.Open("sqlite3", config.DbFile)
	if err != nil {
		return Database{}, err
	}
	if err = db.Ping(); err != nil {
		return Database{}, err
	}

	stmt := `CREATE TABLE IF NOT EXISTS user (
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
		return Database{}, err
	}

	database.Handle = db
	return *database, nil
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
