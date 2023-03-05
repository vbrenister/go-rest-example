package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vbrenister/go-rest-example/api/handlers"
)

const (
	port = ":3000"
)


func main() {
	for path, handler := range handlers.Handlers {
		http.HandleFunc(path, handler)
	}
	
	fmt.Println("Server is running at localhost:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}