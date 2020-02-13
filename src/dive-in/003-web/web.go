package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

var times = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		fmt.Fprintf(w, "Go!")

	} else {
		fmt.Fprintf(w, "Go %s!", r.URL.Path[1:])
	}
	printMemUsage()
	times++
	fmt.Printf("%d", times)
}

// printMemUsage uses 	https://golang.org/pkg/runtime/#MemStats
// to log mem usage.
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	http.HandleFunc("/", handler)
	port := "8081"
	log.Println("Starting at port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
