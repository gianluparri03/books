package data

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"strings"
)

// db is the database connection.
var db *sqlx.DB

// schema is the database schema.
//
//go:embed schema.sql
var schema string

// InitDB initializes the sqlite database connection and makes sure the schema
// is created.
func InitDB(path string) (err error) {
	// Opens the connection
	db, err = sqlx.Open("sqlite", path)
	if err != nil {
		return err
	}

	// Executes each statement of the schema
	for _, q := range strings.Split(schema, ";") {
		if q != "" {
			if _, err = db.Exec(q + ";"); err != nil {
				return err
			}
		}
	}

	return nil
}
