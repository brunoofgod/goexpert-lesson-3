package usecase

import "github.com/brunoofgod/goexpert-lesson-3/internal/entity"

type ListOrderInputDTO struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListOrderOutputDTO struct {
	Orders      []OrderOutputDTO `json:"orders"`
	CurrentPage int              `json:"currentPage"`
	TotalPages  int              `json:"totalPages"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute(input ListOrderInputDTO) (*ListOrderOutputDTO, error) {
	orders, total, err := l.OrderRepository.List(input.Page, input.Limit)

	if err != nil {
		return nil, err
	}

	ordersOutput := []OrderOutputDTO{}
	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		},
		)
	}

	listOrderInputDTO := ListOrderOutputDTO{
		Orders:      ordersOutput,
		CurrentPage: input.Page,
		TotalPages:  total,
	}

	return &listOrderInputDTO, nil
}
