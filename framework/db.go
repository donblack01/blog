package framework

import (
	"blog/config"
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

var Db *sql.DB

func Instance() *sql.DB {
	once.Do(func() {
		dataSource := config.Db_username + ":" + config.Db_password +
			"@" + config.Db_linkType + "(" + config.Db_host + ":" + config.Db_port + ")/" + config.Db_database + "?charset=utf8mb4"
		Db, _ = sql.Open(config.Db_type, dataSource)
	})
	return Db
}
