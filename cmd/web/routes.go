// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func(app *application) routes() http.Handler {

	mux := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	mux.HandlerFunc(http.MethodGet, "/", app.home)
	mux.HandlerFunc(http.MethodGet, "/about", app.about)
	mux.HandlerFunc(http.MethodGet, "/create", app.create)

	return mux
}