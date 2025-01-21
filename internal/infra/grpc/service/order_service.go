package service

import (
	"context"

	"github.com/brunoofgod/goexpert-lesson-3/internal/infra/grpc/pb"
	"github.com/brunoofgod/goexpert-lesson-3/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	// ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
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

// func (s *OrderService) ListOrder(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
// 	output, err := s.ListOrderUseCase.Execute(usecase.ListOrderInputDTO{
// 		Page:  int(in.Page),
// 		Limit: int(in.Limit),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ListOrderResponse{
// 		Orders: []*pb.ListOrderResponse_Order{},
// 	}, nil
// }
