package handler

type Handlers struct {
	User    *UserHandler
	Product *ProductHandler
}

type HandlersParam struct {
	User    *UserHandler
	Product *ProductHandler
}

func NewHandlers(h HandlersParam) *Handlers {
	return &Handlers{
		User:    h.User,
		Product: h.Product,
	}
}
