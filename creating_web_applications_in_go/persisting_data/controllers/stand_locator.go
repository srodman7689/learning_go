package controllers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/controllers/util"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/viewmodels"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content Type", "application/json")
	data, err := json.Marshal(vm)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}
