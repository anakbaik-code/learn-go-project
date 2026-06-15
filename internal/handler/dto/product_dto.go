package dto

type ProductResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
type ProductRequest struct {
	Name  string
	Price int32
}

type ProductUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Price int32  `json:"price" vaalidate:"required,gt=0"`
}
