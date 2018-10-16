package main

import (
	"log"
	"net/http"

	"github.com/steven-liao/objectserver/pkg/http/rest"
)

func main() {
	rest.Init()
	r := rest.Handler()
	if r == nil {
		log.Fatalln("Invalid http handler")
	}

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("http error")
	}
}
