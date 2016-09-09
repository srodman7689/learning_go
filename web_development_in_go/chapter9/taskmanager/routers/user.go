package routers

import (
	"github.com/gorilla/mux"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
