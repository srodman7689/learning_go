package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/common"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/routers"
)

//Entry point of the program
func main() {
	// Calls startup logic
	common.StartUp()

	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
