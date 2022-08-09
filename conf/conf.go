package conf

import (
	"go-blog/model"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbPort     string
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
)

func Init() {
	// TODO: read conf.init file
	HttpPort = ":8080"
	DbUser = "root"
	DbPassword = "123456"
	DbHost = "127.0.0.1"
	DbPort = "3306"
	DbName = "go-blog"

	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	// TODO: cache redis
}
