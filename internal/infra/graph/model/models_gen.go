// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"Price"`
	Tax        float64 `json:"Tax"`
	FinalPrice float64 `json:"FinalPrice"`
}

type OrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"Price"`
	Tax   float64 `json:"Tax"`
}

type OrderPagination struct {
	Orders      []*Order `json:"orders"`
	CurrentPage int      `json:"currentPage"`
	TotalPages  int      `json:"totalPages"`
}

type Query struct {
}
