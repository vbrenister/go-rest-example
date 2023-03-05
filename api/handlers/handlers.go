package handlers

import (
	"fmt"
	"net/http"
)


var Handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/": HelloWorld,
}

func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello world\n")
}