package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/start", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Print("Starting server at port 8082\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/start" {
		http.Error(w, "404 not found bruh!", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Failed to parse form %v", err)
		return
	}

	fmt.Fprintf(w, "POST request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
