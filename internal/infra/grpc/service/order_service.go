package service

import (
	"context"

	"github.com/kleytonsolinho/golang-clean-arch/internal/infra/grpc/pb"
	"github.com/kleytonsolinho/golang-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	OrderListUseCase 	 usecase.FindAllOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, orderListUseCase usecase.FindAllOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		OrderListUseCase: orderListUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	output, err := s.OrderListUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderListResponse []*pb.CreateOrderResponse

	for _, order := range output {
		orderListResponse = append(orderListResponse, &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.OrderList{
		Orders: orderListResponse,
	}, nil
}
