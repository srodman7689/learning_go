package main

import (
	"net/http"

	"github.com/srodman7689/learning_go/web_development_in_go/chapter10/httptestbdd/lib"
)

func main() {
	routers := lib.SetUserRoutes()
	http.ListenAndServe(":8080", routers)
}
