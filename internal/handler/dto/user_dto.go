package dto

type AddressRequest struct {
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type CreateUserRequest struct {
	Name      string           `json:"name" validate:"required,min=3,max=50"`
	Email     string           `json:"email" validate:"required,email"`
	Addresses []AddressRequest `json:"addresses" validate:"required,dive"` 
}

type UpdateUserRequest struct {
	Name      string           `json:"name" validate:"required,min=3"`
	Email     string           `json:"email" validate:"required,email"`
	Addresses []AddressRequest `json:"addresses" validate:"required,dive"`
}

type UploadAvatar struct {
	AvatarUrl string `json:"avatar_url"`
}

type AddressResponse struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}
type UserResponse struct {
	ID        int64             `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	AvatarUrl string            `json:"avatar_url"`
	Addresses []AddressResponse `json:"addresses"`
}
