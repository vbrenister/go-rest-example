package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vbrenister/go-rest-example/pkg/db/products"
)

type Handler struct {
	Routes map[string]func(http.ResponseWriter, *http.Request)

	products products.Repo
}

func NewHandler(products products.Repo) *Handler {
	handler := &Handler{
		products: products,
	}

	attachHandlers(handler)

	return handler
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Products store\n")
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	prs, err := h.products.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prs)
}

func attachHandlers(h *Handler) {
	h.Routes = map[string]func(http.ResponseWriter, *http.Request){
		"/":         h.home,
		"/products": h.getProducts,
	}
}
