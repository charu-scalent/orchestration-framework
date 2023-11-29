package main

import (
	"context"
	"fmt"
)

type OrderInterface interface {
	CreateOrder(ctx context.Context) (interface{}, error)
	RollbackCreateOrder(ctx context.Context) error
}

type Order struct {
	Id    int
	Total float64
}

func (o Order) CreateOrder(ctx context.Context) (interface{}, error) {

	fmt.Printf("Order Instance - Create order executed \n")

	order := Order{Id: 1, Total: 200}
	return order, nil
}

func (o Order) RollbackCreateOrder(ctx context.Context) error {

	fmt.Printf("Order Instance - Create order rolled back \n")
	return nil
}
