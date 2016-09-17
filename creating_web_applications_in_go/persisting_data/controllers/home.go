package controllers

import (
	"net/http"
	"text/template"

	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/controllers/util"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/viewmodels"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")

	vm := viewmodels.GetLogin()

	this.loginTemplate.Execute(responseWriter, vm)
}
