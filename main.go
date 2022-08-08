package main

import (
	"go-blog/conf"
	"go-blog/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
