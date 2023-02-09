package main

import (
	"net/http"

	"github.com/Robertoangeldom/poll/helpers"
)

// handler fuction
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	helpers.Rendertmpl(w, "./Static/html/home.page.tmpl")

}

// handler fuction
func (app *application) About(w http.ResponseWriter, r *http.Request) {
	helpers.Rendertmpl(w, "./Static/html/about.page.tmpl")
}

// handler fuction
func (app *application) HandlerPoll(w http.ResponseWriter, r *http.Request) {
	helpers.Rendertmpl(w, "./Static/html/create.poll.tmpl")
}
