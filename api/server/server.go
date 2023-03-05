package server

import (
	"net/http"

	"github.com/vbrenister/go-rest-example/api/handlers"
)

func Run(addr string) error {
	for path, handler := range handlers.Handlers {
		http.HandleFunc(path, handler)
	}

	return http.ListenAndServe(addr, nil)
}
