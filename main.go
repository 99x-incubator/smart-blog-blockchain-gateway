package main

import (
	"github.com/Blogchain-Gateway/api/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":6969", router))

}
