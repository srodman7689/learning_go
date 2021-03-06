package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/srodman7689/learning_go/creating_web_applications_in_go/controller_layer_pt1/controllers"
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
