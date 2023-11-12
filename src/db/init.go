package db

import (
	"database/sql"
	"sync"
)

type DB struct {
	Mutex sync.Mutex
	Conn  *sql.DB
}

func InitDB() *DB {
	return initSQLite(tablesSQLite)
}
