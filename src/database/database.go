package database

import (
	"database/sql"
	"fmt"

	"github.com/mahiro72/go-mysql-ca/config"

	_ "github.com/go-sql-driver/mysql"
)

// NewDBはMySQLを接続し、sql.DBオブジェクトのポインタを返します
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL : %w", err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to db ping : %w", err)
	}

	return db, nil
}
