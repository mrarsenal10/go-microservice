package main

type service struct {
	store OrderService
}

func NewService(store OrderService) *service {
	return &service{store: store}
}

func(s *service) CreateOrder(ctx context.Context) error {
	return s.store.Create(ctx)
}