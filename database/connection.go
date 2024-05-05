package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"sync"
)

var (
	dbConn *pgx.Conn
	once   sync.Once
)

func initDBConnection() error {
	connString := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %v", err)
	}
	dbConn = conn
	return nil
}

// GetDBConnection initializes and returns the database connection
func GetDBConnection() (*pgx.Conn, error) {
	var err error
	once.Do(func() {
		err = initDBConnection()
	})
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// CloseDBConnection closes the database connection
func CloseDBConnection() error {
	if dbConn != nil {
		dbConn.Close(context.Background())
	}
	return nil
}
