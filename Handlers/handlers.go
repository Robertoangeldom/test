package handlers

import (
	"net/http"

	"github.com/Robertoangeldom/poll/helpers"
)

// handler fuction
func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	helpers.Rendertmpl(w, "./Static/html/home.page.tmpl")

}

// handler fuction
func About(w http.ResponseWriter, r *http.Request) {
	helpers.Rendertmpl(w, "./Static/html/about.page.tmpl")
}

// handler fuction
func HandlerPoll(w http.ResponseWriter, r *http.Request) {
	helpers.Rendertmpl(w, "./Static/html/create.poll.tmpl")
}


