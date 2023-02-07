package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("welcome to home page"))

	ts, err := template.ParseFiles("./Static/html/home.page.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to home page"))
}

func HandlerPoll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to home page"))
}
