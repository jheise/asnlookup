package main

import (
	// standard
	"flag"
	"net/http"
	"os"

	// external
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// config options
	address  string
	port     string
	addrport string

	// channel
	requests chan *ASNRequest
)

func init() {
	flag.StringVar(&address, "address", "0.0.0.0", "address to bind to")
	flag.StringVar(&port, "port", "9999", "port to bind to")
	flag.Parse()

	addrport = address + ":" + port
}

func main() {
	requests = make(chan *ASNRequest)
	defer close(requests)

	// create and start new processor
	processor := NewASNProcessor(requests)
	go processor.Process()

	router := mux.NewRouter()
	router.HandleFunc("/v1/asn/{addr}", ASNHandler).Methods("GET")
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	http.ListenAndServe(addrport, loggedRouter)
}
