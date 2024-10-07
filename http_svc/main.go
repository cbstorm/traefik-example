package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var host string
var port int
var sName string

func main() {
	flag.StringVar(&host, "host", "", "Host")
	flag.IntVar(&port, "port", 0, "Port")
	flag.StringVar(&sName, "s_name", "", "Service name")
	flag.Parse()
	if port == 0 {
		log.Fatalf("Port is required")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello, Welcome to server %s", sName)))
	})
	log.Printf("HTTP server [%s] listening at %s:%d", sName, host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux); err != nil {
		log.Fatalf("Could not listen due to an error: %v", err)
	} else {
		log.Println("HTTP Service listening")
	}
}
