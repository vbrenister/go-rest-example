package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":3000"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is running at localhost:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}