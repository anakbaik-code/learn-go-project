package dto

type ProductResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int32  `json:"price"`
	IsActive  bool   `json:"is_active"`
	SalePrice int32  `json:"sale_price"`
	Sku       string `json:"sku"`
}

// nested struct validation
type DiscountDetail struct {
	IsActive  bool  `json:"is_active" validate:"required"`
	SalePrice int64 `json:"sale_price" validate:"required_if=IsActive true,gt=0"`
}
type CreateProductNestedRequest struct {
	Name     string         `json:"name" validate:"required,min=3"`
	Price    int64          `json:"price" validate:"required,gt=0"`
	Discount DiscountDetail `json:"discount" validate:"required"`
	Sku      string         `json:"sku" validate:"required,valid_sku"`
}

type ProductUpdateRequest struct {
	Name      string `json:"name" validate:"required,min=3,max=100"`
	Price     int32  `json:"price" validate:"required,gt=0"`
	IsActive  bool   `json:"is_active" validate:"required"`
	SalePrice int32  `json:"sale_price" validate:"required"`
	Sku       string `json:"sku" validate:"required,valid_sku"`
}
