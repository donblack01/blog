package framework

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var once sync.Once

var Db *sql.DB

func Instance() *sql.DB {
	once.Do(func() {
		Db, _ = sql.Open("mysql", "root:123456@tcp(65.49.214.22:3306)/blog?charset=utf8mb4")
	})
	return Db
}
