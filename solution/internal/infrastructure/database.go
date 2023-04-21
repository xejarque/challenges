package infrastructure

import (
	"database/sql"
	"fmt"
)

type SqlDB struct {
	db *sql.DB
}

func NewSqlDB(config *DatabaseConfig) (*SqlDB, error) {
	// Create the connection string using the configuration details
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DatabaseName)

	// Open a new database connection
	db, err := sql.Open(config.Driver, connectionString)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Return the database adapter instance
	return &SqlDB{db: db}, nil
}

func (p *SqlDB) Close() error {
	// Close the database connection
	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func (p *SqlDB) Execute(query string, args ...interface{}) (int64, error) {
	// Execute a database command that doesn't return any rows
	result, err := p.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected by the command
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
