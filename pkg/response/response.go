package response

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewSuccessResponse[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}
