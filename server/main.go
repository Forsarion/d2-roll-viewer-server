package main

import (
	"server/manifest"
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
