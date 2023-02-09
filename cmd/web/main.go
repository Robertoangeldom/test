package main

import (
	"flag"
	"log"
	"net/http"
)

//create a new type
type application struct{}

func main() {

	//create a flag for speccifing the port number
	//when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	flag.Parse()

	//create an instants of application type 
	app := &application{}

	//create multiplexer
	mux := http.NewServeMux()

	//create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/handlerPoll", app.HandlerPoll)

	log.Printf("starting server on port %s", *addr)
	err := http.ListenAndServe( *addr, mux)
	log.Fatal(err)
}
