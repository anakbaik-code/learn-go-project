package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/handler/dto"
	"go-dbsqlc/internal/service"
	validate "go-dbsqlc/internal/validator"
	"go-dbsqlc/pkg/response"
	"log/slog"
	"net/http"
	"strconv"
)

type UserHandler struct {
	validator *validator.Validate
	service   service.UserService
	log       *slog.Logger
}

func NewUserHandler(v *validator.Validate, l *slog.Logger, s service.UserService) *UserHandler {
	return &UserHandler{
		validator: v,
		service:   s,
		log:       l.With("component", "users_handler"),
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	//  decode JSON request
	var req dto.CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body json", http.StatusBadRequest)
		return
	}
	var domainAddress []domain.Address
	for _, item := range req.Addresses {
		domainAddress = append(domainAddress, domain.Address{
			Street:  item.Street,
			City:    item.City,
			Country: item.Country,
		})
	}

	// mapping ke domain
	user := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Addresses: domainAddress,
	}

	// validator
	if err := validate.ValidateCreateUser(h.validator, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// panggil service
	result, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		h.log.Error("failed create user from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var addressesResponse []dto.AddressResponse
	for _, addr := range result.Addresses {
		addressesResponse = append(addressesResponse, dto.AddressResponse{
			Street:  addr.Street,
			City:    addr.City,
			Country: addr.Country,
		})
	}

	// Mapping Response
	userResponse := dto.UserResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		Addresses: addressesResponse,
	}

	finalResponse := response.NewSuccessResponse(
		"User Created", userResponse,
	)

	// response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if idStr == "" {
		http.Error(w, "id must fiil", http.StatusBadRequest)
		return
	}

	// convert string to int
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id invalid", http.StatusBadRequest)
		return
	}

	// validator
	if err := validate.ValidateGetUserByID(h.validator, id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// panggil service
	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		h.log.Error("failed to get product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var addressesResponse []dto.AddressResponse
	for _, addr := range user.Addresses {
		addressesResponse = append(addressesResponse, dto.AddressResponse{
			Street:  addr.Street,
			City:    addr.City,
			Country: addr.Country,
		})
	}

	// Mapping Response
	userResponse := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Addresses: addressesResponse,
	}
	finalResponse := response.NewSuccessResponse(
		"succesfully get user id ",
		userResponse,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)

}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	usersDomain, err := h.service.ListUsers(r.Context())
	if err != nil {
		h.log.Error("failed to get list product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userResponse []dto.UserResponse
	for _, user := range usersDomain {
		var addressesResponse []dto.AddressResponse
		for _, addr := range user.Addresses {
			addressesResponse = append(addressesResponse, dto.AddressResponse{
				Street:  addr.Street,
				City:    addr.City,
				Country: addr.Country,
			})
		}
		userResponse = append(userResponse, dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Addresses: addressesResponse,
		})
	}
	// Mapping DTO use Slice

	finalResponse := response.NewSuccessResponse(
		"Successfully fetched user list",
		userResponse,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req dto.UpdateUserRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	// validator
	if err := validate.ValidateUpdateUser(h.validator, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// mapping from DTO to slice
	var domainAddresses []domain.Address
	for _, item := range req.Addresses {
		domainAddresses = append(domainAddresses, domain.Address{
			Street:  item.Street,
			City:    item.City,
			Country: item.Country,
		})
	}

	// add domainAddress to object domain.User
	userDomain := domain.User{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Addresses: domainAddresses,
	}

	err = h.service.UpdateUser(r.Context(), userDomain)
	if err != nil {
		h.log.Error("failed update product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var addressesResponse []dto.AddressResponse
	for _, addr := range userDomain.Addresses {
		addressesResponse = append(addressesResponse, dto.AddressResponse{
			Street:  addr.Street,
			City:    addr.City,
			Country: addr.Country,
		})
	}

	userResponse := dto.UserResponse{
		ID:        id,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Addresses: addressesResponse,
	}

	finalResponse := response.NewSuccessResponse[any](
		"User updated successfully", userResponse,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(finalResponse)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(r.Context(), id)
	if err != nil {
		h.log.Error("failed delete product from service", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalResponse := response.NewSuccessResponse[any](
		"user deleted successfully", nil,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(finalResponse)
}

func (h *UserHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")

	if err != nil {
		http.Error(w, "failed upload", http.StatusBadRequest)
		return
	}

	// validator
	if err := validate.ValidateImage(h.validator, file, header); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	url, err := h.service.UploadAvatar(r.Context(), id, file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping DTO avatarUrl
	avatarResponse := dto.UploadAvatar{
		AvatarUrl: url,
	}
	finalResponse := response.NewSuccessResponse(
		"Avatar Uploaded successfully",
		avatarResponse,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResponse)

}
