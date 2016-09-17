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

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		product, err := models.GetProductById(id)

		if err == nil {

			w.Header().Add("Content Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()

			vm := viewmodels.GetProduct(product.Name())
			vm.Product = converters.ConvertProductToViewModel(product)

			this.template.Execute(responseWriter, vm)
		}
	} else {
		w.WriteHeader(404)
	}
}
