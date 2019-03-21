package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxIdleConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to MySQL,err:" + err.Error())
		os.Exit(1)
	}
}

// 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
