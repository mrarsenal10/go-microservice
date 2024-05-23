type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.createOrder)
}

func (h *handler) createOrder(w http.ResponseWriter, r *http.Request) {
	// create order
}