package db

import (
	"database/sql"
	"path/filepath"
	"trl/conf"
	"trl/srv/errs"
	"trl/srv/logger"

	"github.com/huandu/go-sqlbuilder"
	_ "github.com/mattn/go-sqlite3"
)

func initSQLite(tablesForInit []string) *DB {
	filePath := filepath.Join(conf.Settings.Db.Path, conf.Settings.Db.FileName)
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		errs.LogFatalIfError(err)
	}
	for _, table := range tablesForInit {
		if _, err := db.Exec(table); err != nil {
			errs.LogFatalIfError(err)
		}
	}
	logger.Debug("SQLite database has been initialized")
	return &DB{Conn: db}
}

const WordsTableName = "words"

var createWordsTableSQLite = sqlbuilder.
	CreateTable(WordsTableName).
	IfNotExists().
	Define("id", "INTEGER", "NOT NULL", "PRIMARY KEY").
	Define("word", "TEXT", "NOT NULL").
	Define("translation", "TEXT", "NOT NULL").
	Define("transcription", "TEXT", "NOT NULL").
	Define("created_at", "TEXT", "NOT NULL").
	Define("updated_at", "TEXT", "NOT NULL").
	String()

var tablesSQLite = []string{
	createWordsTableSQLite,
}
