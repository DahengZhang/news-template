package db

import (
	"dahengzhang/news/config"
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
	dbConn, err = sql.Open("mysql", config.Conf.Mysql.User+":"+config.Conf.Mysql.Pwd+"@tcp("+config.Conf.Mysql.Host+":"+config.Conf.Mysql.Port+")/"+config.Conf.Mysql.Database+"?charset=utf8&multiStatements=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	// 创建数据表
	_, err = dbConn.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal(err.Error())
	}
}
