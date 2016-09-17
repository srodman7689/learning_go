package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/controller_layer_pt1/controllers/util"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/controller_layer_pt1/viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProducts(id)

		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}
