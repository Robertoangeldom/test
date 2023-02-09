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

	//create an of application type 
	app := &application{}
	//create a customizable server
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),

	}
	log.Printf("starting server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
