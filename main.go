package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHendler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n\n")

	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Age = %s\n", age)
}

func helloHendler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHendler)
	http.HandleFunc("/hello", helloHendler)

	fmt.Printf("Server Stated at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
