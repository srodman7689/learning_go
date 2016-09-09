package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/common"
	"github.com/srodman7689/learning_go/web_development_in_go/chapter9/taskmanager/controllers"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	noteRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteById).Methods("GET")
	noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	noteRouter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))
	return router
}
