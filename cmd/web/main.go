package main

import (
	"log"
	"net/http"

	handlers "github.com/Robertoangeldom/poll/Handlers"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/handlerPoll", handlers.HandlerPoll)

	log.Print("starting server on port:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
