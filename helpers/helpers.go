package helpers

import (
	"html/template"
	"log"
	"net/http"
)

func Rendertmpl(w http.ResponseWriter, tmpl string) {
	//allows you to see the webpage
	ts, err := template.ParseFiles(tmpl)

	//checks for errors
	if err != nil {
		log.Print(err.Error())
		//sends the error to the user
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
