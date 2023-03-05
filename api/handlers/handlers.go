package handlers

import (
	"fmt"
	"net/http"

	"github.com/vbrenister/go-rest-example/pkg/db/products"
)

type Handler struct {
	Handlers map[string]func(http.ResponseWriter, *http.Request)

	products products.Repo
}

func NewHandler(products products.Repo) *Handler {
	handler := &Handler{
		products: products,
	}

	attachHandlers(handler)

	return handler
}

func attachHandlers(h *Handler) {
	h.Handlers = map[string]func(http.ResponseWriter, *http.Request){
		"/": h.HelloWorld,
	}
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(h.products.GetAll())
	fmt.Fprintf(w, "Hello world\n")
}
