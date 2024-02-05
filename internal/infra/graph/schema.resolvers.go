package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	model1 "github.com/kleytonsolinho/golang-clean-arch/graph/model"
	"github.com/kleytonsolinho/golang-clean-arch/internal/infra/graph/model"
	"github.com/kleytonsolinho/golang-clean-arch/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model1.OrderInput) (*model1.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model1.Order, error) {
	output, err := r.OrderListUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*model1.Order
	for _, order := range output {
		orders = append(orders, &model1.Order{
			ID:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) orders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.OrderListUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for _, order := range output {
		orders = append(orders, &model.Order{
			ID:         order.ID,
			Price:      float64(order.Price),
			Tax:        float64(order.Tax),
			FinalPrice: float64(order.FinalPrice),
		})
	}
	return orders, nil
}
