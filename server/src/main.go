package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/manifest", manifest)
	http.HandleFunc("/findHelmets", findHelmets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
