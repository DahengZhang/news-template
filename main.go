package main

import (
	"dahengzhang/news/config"
	_ "dahengzhang/news/config"
	"dahengzhang/news/router"
	"net/http"
	"strconv"
)

func main() {
	r := router.SetRouter()
	http.ListenAndServe(":"+strconv.Itoa(config.Conf.Server.Port), r)
}
