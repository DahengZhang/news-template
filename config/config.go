package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Server 服务端配置
type server struct {
	Port int `json:"port"`
}

// Mysql 数据库配置
type mysql struct {
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

// Config 配置信息
type Config struct {
	Server server `json:"server"`
	Mysql  mysql  `json:"mysql"`
}

// Conf 全局配置变量
var Conf Config

func init() {
	viper.AddConfigPath("./config") // 第一个搜索路径
	if os.Getenv("NEWS_ENV") == "dev" {
		viper.SetConfigName("dev") // 设置配置文件名 (不带后缀)
	} else {
		viper.SetConfigName("pro")
	}
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		log.Fatal(err.Error())
	}
	viper.Unmarshal(&Conf)
}
