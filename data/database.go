package data

import (
	"github.com/jmoiron/sqlx"
	"io"
	_ "modernc.org/sqlite"
	"os"
	"strings"
)

var db *sqlx.DB

// InitDB initializes the sqlite database connection and makes sure the schema
// is created.
func InitDB(path string) (err error) {
	// Opens the connection
	db, err = sqlx.Open("sqlite", path)
	if err != nil {
		return err
	}

	// Opens the schema file
	f, err := os.Open("data/schema.sql")
	defer f.Close()
	if err != nil {
		return err
	}

	// Reads the schema file
	s, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// Executes each statement
	for _, q := range strings.Split(string(s), ";") {
		if q != "" {
			if _, err = db.Exec(q + ";"); err != nil {
				return err
			}
		}
	}

	return nil
}
