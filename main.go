package main

import (
	"fmt"
	"log"
	"net/http"
)

// Port constant for service
const Port = "80"

func main() {
	http.HandleFunc("/", sendGopher)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

func sendGopher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DoT API version 1.0\n")
}

