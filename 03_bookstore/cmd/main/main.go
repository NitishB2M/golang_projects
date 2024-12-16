package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitishb2m/golang_projects/03_bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
