package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type MysqlCfg struct {
	Dsn    string
	Db     string
	Tables []string
}

func NewDB(cfg *MysqlCfg) {
	dbc, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		panic(err)
	}
	if err = dbc.Ping(); err != nil {
		fmt.Printf("db ping err %v", err)
		panic(err)
	} else {
		db = dbc
	}
}

func Ins() *sql.DB {
	return db
}
