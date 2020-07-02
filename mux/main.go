package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	hostname string
	port     int
)

/* register command line options */
func init() {
	flag.StringVar(&hostname, "hostname", "0.0.0.0", "The hostname or IP on which the REST server will listen")
	flag.IntVar(&port, "port", 8080, "The port on which the REST server will listen")
}

func main() {
	flag.Parse()
	var address = fmt.Sprintf("%s:%d", hostname, port)
	log.Println("REST service listening on", address)

	// register router
	router := mux.NewRouter().StrictSlash(true)
	router.
		HandleFunc("/api/service/get", MyGetHandler).
		Methods("GET")
	router.
		HandleFunc("/api/service/post", MyPostHandler).
		Methods("POST")

	// start server listening
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}

	log.Println("Server end")
}
