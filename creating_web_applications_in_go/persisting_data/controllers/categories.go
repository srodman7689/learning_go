package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/controllers/util"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/converters"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/models"
	"github.com/srodman7689/learning_go/creating_web_applications_in_go/persisting_data/viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()

	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoyToViewModel(category))
	}

	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM

	w.Header().Add("Content Type", "text/html")
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
		category, err := models.GetCategoryById(id)

		if err == nil {

			w.Header().Add("Content Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()

			vm := viewmodels.GetProducts(category.Title())
			productVMs := []viewmodels.Product{}

			for _, product := range category.Products() {
				productVMs = append(productVMs, converters.ConvertProductToViewModel(product))
			}

			vm.Products = productVMs

			this.template.Execute(responseWriter, vm)
		}
	} else {
		w.WriteHeader(404)
	}
}
