package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Database interface
type Database interface {
	Connect()
	Disconnect()
	Query(query string) (*sql.Rows, error)
}

// Database connection "manager" main struct, holds the connection globally
type DatabaseImpl struct {
	Connection *sql.DB
	Server     string
	User       string
	Password   string
	Database   string
}

const (
	NOCONNECTION = "Unable to connecto to database: %s"
)

// Connect connects to database
func (c *DatabaseImpl) Connect() {
	var err error

	// Start database connection
	c.Connection, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Server, c.Database))

	if err != nil {
		panic(err.Error()) // Service should not start if not connected to database
	}

	// Validate DSN data:
	err = c.Connection.Ping()
	if err != nil {
		panic(err.Error()) // Service should not start if not connected to database
	}

	// Additional settings, those should be taken from a config file.
	c.Connection.SetConnMaxLifetime(time.Minute * 3)
	c.Connection.SetMaxOpenConns(10)
	c.Connection.SetMaxIdleConns(10)
}

// Disconnect from the database
func (c DatabaseImpl) Disconnect() {
	c.Connection.Close()
}

// Query launches a query against the database
func (c DatabaseImpl) Query(query string) (*sql.Rows, error) {
	// Prepare statement for reading data
	results, err := c.Connection.Query(query)
	if err != nil {
		log.Printf(err.Error()) // Error is logged for debug
		return nil, NoConnectionError(err.Error())
	}
	return results, nil
}

//NoConnectionError returns a generic error for database
func NoConnectionError(message string) error {
	return fmt.Errorf(NOCONNECTION, message)
}
