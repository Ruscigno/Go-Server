package main

import (
	// "fmt"
	// "io"
	"fmt"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey, here is your request: %s", req.URL.Path)
	log.Printf("Received request for path: %s", req.URL.Path)
}

func main() {
	log.Printf("Starting web server on 8180")
	handler := http.HandlerFunc(helloServer)
	err := http.ListenAndServe(":8180", handler)
	if err != nil {
		log.Fatal("Could not listen on port: 8180", err)
	}
}
