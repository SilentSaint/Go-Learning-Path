package main

import (
	"fmt"
	"log"
	"net/http"
)

// this server version is a minimal echo server
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH=%q\n", r.URL.Path)
}
