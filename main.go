package main

import (
	"food-app/presentation"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(presentation.Serve)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
