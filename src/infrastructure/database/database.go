package database

import (
	"fmt"

	"github.com/mahiro72/go-mysql-ca/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnはDBとの接続情報をもつ構造体です
type Conn struct {
	DB *sqlx.DB
}

// NewConnはMySQLを接続し、sql.DBオブジェクトのポインタをもつ構造体を返します
func NewConn() (*Conn, error) {
	db, err := sqlx.Connect("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL : %w", err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to db ping : %w", err)
	}

	return &Conn{DB: db}, nil
}