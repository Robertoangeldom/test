package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	//create multiplexer
	mux := http.NewServeMux()

	//create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/handlerPoll", app.HandlerPoll)

	return mux
}