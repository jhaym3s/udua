package product

import (
	"log"
	"net/http"

	jsn "github.com/jhaym3s/udua/internal/json"
)

type handler struct {
	Service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		Service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.Service.ListProducts(r.Context())
	if err != nil {
		log.Println("Failed to list products:", err)
		jsn.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to list products"})
		return
	} 
	products := struct {
		Produdcts []string `json:"products"`
	}{
		Produdcts: []string{"Product 1", "Product 4", "Product 3"},
	}

	jsn.WriteJSON(w, http.StatusOK, products)

}
