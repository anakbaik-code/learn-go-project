package handler

import (
	"encoding/json"
	"go-dbsqlc/internal/service"
	"log/slog"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service service.ProductService
	log     *slog.Logger
}

func NewProductHandler(logger *slog.Logger, s service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
		log:     logger.With("component", "product_handler"),
	}
}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "id must fill", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id Ivalid", http.StatusBadRequest)
		return
	}

	h.log.Debug("fetching product details", "product_id", id)

	product, err := h.service.GetProduct(r.Context(), id)
	if err != nil {
		h.log.Error("failed to get product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
