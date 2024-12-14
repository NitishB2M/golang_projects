package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(rw, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(rw, "Name: %v\nAddress: %v", name, address)
}

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(rw, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
