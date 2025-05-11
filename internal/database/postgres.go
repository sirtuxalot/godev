/* internal/database/postgres.go */

package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

// holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifetime = 5 * time.Minute

// creates database pool for postgresql
func ConnectSQL(connectString string) (*DB, error) {
	appDB, err := NewDatabase(connectString)
	if err != nil {
		panic(err)
	}
	appDB.SetMaxIdleConns(maxIdleDBConn)
	appDB.SetMaxOpenConns(maxOpenDBConn)
	appDB.SetConnMaxIdleTime(maxDBLifetime)
	dbConn.SQL = appDB
	err = testDB(appDB)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// attempts to ping connected database
func testDB(appDB *sql.DB) error {
	err := appDB.Ping()
	if err != nil {
		return err
	}
	return nil
}

// creates a new database for the application
func NewDatabase(connectString string) (*sql.DB, error) {
	appDB, err := sql.Open("pgx", connectString)
	if err != nil {
		return nil, err
	}
	if err = appDB.Ping(); err != nil {
		return nil, err
	}
	return appDB, nil
}
