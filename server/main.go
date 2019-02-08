package main

import (
	manifest "github.com/Berk0ld/d2-roll-viewer-server/server/manifest"
)

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/manifest", manifest.HTTPHandler)
	http.HandleFunc("/findHelmets", FindHelmets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
