package framework

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func OpenDatabase(config *Config) (*sql.DB, error) {
	var database strings.Builder

	//
	// "user:password@tcp(127.0.0.1:3306)/database_name")
	fmt.Fprintf(&database, "%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName)

	db, err := sql.Open("mysql", database.String())
	return db, err
}

func CloseDatabase(db sql.DB) {
	db.Close()
}
