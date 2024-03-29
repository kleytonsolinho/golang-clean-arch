//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/kleytonsolinho/golang-clean-arch/internal/entity"
	"github.com/kleytonsolinho/golang-clean-arch/internal/event"
	"github.com/kleytonsolinho/golang-clean-arch/internal/infra/database"
	"github.com/kleytonsolinho/golang-clean-arch/internal/infra/web"
	"github.com/kleytonsolinho/golang-clean-arch/internal/usecase"
	"github.com/kleytonsolinho/golang-clean-arch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewFindAllOrderUseCase(db *sql.DB) *usecase.FindAllOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewFindAllOrderUseCase,
	)
	return &usecase.FindAllOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
