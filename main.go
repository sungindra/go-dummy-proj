package main

import (
	"dummy/handler"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler.RouteHandler()))
}
