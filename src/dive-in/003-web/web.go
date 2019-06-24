package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		fmt.Fprintf(w, "Go!")

	} else {
		fmt.Fprintf(w, "Go %s!", r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := "8081"
	log.Println("Starting at port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
