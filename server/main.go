package main

import (
	"log"
	"net/http"

	manifest "github.com/Berk0ld/d2-roll-viewer-server/server/manifest"
)

func main() {
	http.HandleFunc("/manifest", manifest.HTTPHandler)
	http.HandleFunc("/findHelmets", FindHelmets)
	port := ":8080"
	log.Printf("Server running at %v!\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
