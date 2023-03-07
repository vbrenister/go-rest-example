package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vbrenister/go-rest-example/pkg/db/products"
)

type Handler struct {
	products products.Repo
}

func New(products products.Repo) *Handler {
	return &Handler{
		products: products,
	}
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
