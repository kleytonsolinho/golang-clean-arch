package graph

import "github.com/kleytonsolinho/golang-clean-arch/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	OrderListUseCase   usecase.FindAllOrderUseCase
}
