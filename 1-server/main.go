package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(write http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(write, "404 - not found", http.StatusNotFound)
		return 
	} 

	if request.Method != "GET" {
		http.Error(write, "method is not supported", http.StatusNotFound)
		return 
	}  

	fmt.Fprintf(write, "hello!")
}

func formHandler(write http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(write, "method is not supported", http.StatusNotFound)
		return 
	}  
	
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(write, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(write, "post request successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(write, "Name - %s\n", name)
	fmt.Fprintf(write, "Address - %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} 
}