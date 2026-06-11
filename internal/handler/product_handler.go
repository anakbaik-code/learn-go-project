package handler

import (
	"encoding/json"
	"go-dbsqlc/internal/handler/dto"
	"go-dbsqlc/internal/service"
	"go-dbsqlc/pkg/response"
	"log/slog"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service service.ProductService
	log     *slog.Logger
}

func NewProductHandler(l *slog.Logger, s service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
		log:     l.With("component", "product_handler"),
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

	// Mapping Response
	productResponse := dto.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	finalResponse := response.NewSuccessResponse(
		"successfully fetched product details",
		productResponse,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}


