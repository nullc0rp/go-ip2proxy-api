package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database connection "manager" main struct, holds the connection globally
type Database struct {
	connection *sql.DB
	Server     string
	User       string
	Password   string
	Database   string
}

const (
	NOCONNECTION = "Unable to connecto to database: %s"
)

// Connect connects to database
func (c *Database) Connect() {
	var err error

	// Start database connection
	c.connection, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Server, c.Database))

	if err != nil {
		panic(err.Error()) // Service should not start if not connected to database
	}

	// Validate DSN data:
	err = c.connection.Ping()
	if err != nil {
		panic(err.Error()) // Service should not start if not connected to database
	}

	// Additional settings, those should be taken from a config file.
	c.connection.SetConnMaxLifetime(time.Minute * 3)
	c.connection.SetMaxOpenConns(10)
	c.connection.SetMaxIdleConns(10)
}

// Disconnect from the database
func (c Database) Disconnect() {
	c.connection.Close()
}

// Query launches a query against the database
func (c Database) Query(query string) (*sql.Rows, error) {
	// Prepare statement for reading data
	results, err := c.connection.Query(query)
	if err != nil {
		log.Printf(err.Error()) // Error is logged for debug
		return nil, NoConnectionError(err.Error())
	}
	return results, nil
}

//New Creates a new database
func (c Database) New(server string, user string, password string, database string) *Database {
	return &Database{
		Server:   server,
		User:     user,
		Password: password,
		Database: database,
	}
}

//NoConnectionError returns a generic error for database
func NoConnectionError(message string) error {
	return fmt.Errorf(NOCONNECTION, message)
}
