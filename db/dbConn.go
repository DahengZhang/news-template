package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	// MYSQL 驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	// 读取 sql 文件
	sqlBytes, err := ioutil.ReadFile("docs/news.sql")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 创建链接池 docker run -d -p 8080:8080 --link news-mysql:db dahengzhang/news
	dbConn, err = sql.Open("mysql", "dahengzhang:000000@tcp(db:3306)/webIM?charset=utf8&multiStatements=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 创建数据表
	_, err = dbConn.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
}
