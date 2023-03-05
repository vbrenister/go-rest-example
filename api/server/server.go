package server

import (
	"log"
	"net/http"

	"github.com/vbrenister/go-rest-example/api/handlers"
	"github.com/vbrenister/go-rest-example/pkg/db/products"
)

func Run(addr string) error {
	prd, err := products.NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	for path, handler := range handlers.NewHandler(prd).Handlers {
		http.HandleFunc(path, handler)
	}

	return http.ListenAndServe(addr, nil)
}
