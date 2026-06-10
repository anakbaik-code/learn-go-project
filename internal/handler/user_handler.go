package handler

import (
	"encoding/json"
	"go-dbsqlc/internal/domain"
	"go-dbsqlc/internal/handler/dto"
	"go-dbsqlc/internal/service"
	"go-dbsqlc/internal/validator"
	"go-dbsqlc/pkg/response"
	"log"
	"net/http"
	"strconv"

	playvalidator "github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service  service.UserService
	validate *playvalidator.Validate
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		service:  s,
		validate: playvalidator.New(),
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	//  decode JSON request
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	log.Printf("REQ: %+v\n", req)

	// mapping ke domain
	user := domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// panggil service
	result, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping Response
	userResponse := dto.UserResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
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

	// panggil service
	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mapping REsponse
	userResponse := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
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
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Mapping DTO use Slice
	userResponse := make([]dto.UserResponse, 0, len(users))
	for _, v := range users {
		userDto := dto.UserResponse{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Name,
		}
		userResponse = append(userResponse, userDto)
	}
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

	err = h.validate.Struct(req)
	if err != nil {
		// Jika validasi gagal, kirim error 400 dan BERHENTI di sini
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// mapping param
	updateParam := domain.UpdateUserParam{
		Name:  req.Name,
		Email: req.Email,
	}

	err = h.service.UpdateUser(r.Context(), id, updateParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalResponse := response.NewSuccessResponse[any](
		"User updated successfully", nil,
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

	err = validator.ValidateImage(file, header)
	if err != nil {
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
