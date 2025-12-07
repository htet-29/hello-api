package main

import (
	"log"
	"net/http"

	"github.com/htet-20/hello-api/handlers/rest"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func main() {
	addr := "localhost:8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("Listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
