package main

import (
	"dahengzhang/news/router"
	"net/http"
)

func main() {
	r := router.SetRouter()
	http.ListenAndServe(":8080", r)
}
