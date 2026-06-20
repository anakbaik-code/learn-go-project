package handler

import (
	"encoding/json"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/handler/dto"
	"go-dbsqlc/internal/service"
	validate "go-dbsqlc/internal/validator"
	"go-dbsqlc/pkg/response"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	validator *validator.Validate
	service   service.ProductService
	log       *slog.Logger
}

func NewProductHandler(v *validator.Validate, l *slog.Logger, s service.ProductService) *ProductHandler {
	return &ProductHandler{
		validator: v,
		service:   s,
		log:       l.With("component", "product_handler"),
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

	// validator
	if err := validate.ValidateProductId(h.validator, id); err != nil {
		http.Error(w, "id invalid", http.StatusBadRequest)
	}

	product, err := h.service.GetProduct(r.Context(), id)
	if err != nil {
		h.log.Error("failed to get product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping Response
	productResponse := dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		IsActive:  product.IsActive,
		SalePrice: product.SalePrice,
	}
	finalResponse := response.NewSuccessResponse(
		"successfully fetched product details",
		productResponse,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductNestedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// validator
	errs := validate.ValidateCreateProduct(h.validator, req)
	
	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		
		// Kirim kumpulan pesan error-nya langsung sebagai JSON object ke Postman
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validasi gagal mase!",
			"errors":  errs, 
		})
		return
	}
	finalSalePrice := req.Discount.SalePrice
	if !req.Discount.IsActive {
		finalSalePrice = 0
	}

	// mapping dto
	product := domain.Product{
		Name:      req.Name,
		Price:     int32(req.Price),
		IsActive:  req.Discount.IsActive,
		SalePrice: int32(finalSalePrice),
		Sku:       req.Sku,
	}

	// service
	result, err := h.service.CreateProduct(r.Context(), product)
	if err != nil {
		h.log.Error("failed create user from service ", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping Response
	userResponse := dto.ProductResponse{
		ID:        result.ID,
		Name:      result.Name,
		Price:     result.Price,
		IsActive:  result.IsActive,
		SalePrice: result.SalePrice,
		Sku:       result.Sku,
	}

	finalResponse := response.NewSuccessResponse(
		"Product Created",
		userResponse,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	productDomain, err := h.service.ListProduct(r.Context())
	if err != nil {
		h.log.Error("failed to get list users from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var productResponse []dto.ProductResponse
	for _, product := range productDomain {
		productResponse = append(productResponse, dto.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			IsActive:  product.IsActive,
			SalePrice: product.SalePrice,
			Sku:       product.Sku,
		})
	}
	finalResponse := response.NewSuccessResponse(
		"successfully fetched products list",
		productResponse,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}
