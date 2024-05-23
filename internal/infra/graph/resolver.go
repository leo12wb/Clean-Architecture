package graph

import "github.com/leo12wb/Clean-Architecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase.CreateOrderUseCase
	usecase.ListOrdersUseCase
}
