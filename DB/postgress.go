package db
import (
	"database/sql"
	"fmt"
	"sync"
	_"github.com/lib/pq"
)

var dbInstance *sql.DB
var dbInstanceError error
var dbOnce sync.Once

const (
	// Update the connection string to match your PostgreSQL configuration
	// Replace "user", "password", "host", "port", and "dbname" with your actual PostgreSQL credentials and database name
	ConnectionStr = "user=postgres password=postgres host=localhost port=5432 dbname=funn sslmode=disable"
)

func GetPostgresDB() (*sql.DB, error) {
	dbOnce.Do(func() {
		db, err := sql.Open("postgres", ConnectionStr)
		if err != nil {
			dbInstanceError = fmt.Errorf("failed to connect to PostgreSQL: %v", err)
			return
		}

		err = db.Ping()
		if err != nil {
			dbInstanceError = fmt.Errorf("failed to ping PostgreSQL: %v", err)
			return
		}

		dbInstance = db

		err = createSchema(db)
		if err != nil {
			dbInstanceError = fmt.Errorf("failed to create schema: %v", err)
			dbInstance = nil
		}
	})
	return dbInstance, dbInstanceError
}

func createSchema(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
		email TEXT NOT NULL,
        password TEXT NOT NULL,
        credit FLOAT DEFAULT 0,
        area TEXT,
        address TEXT
    );
    `
    _, err := db.Exec(query)
    return err
}
