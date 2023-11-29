package main

import (
	"context"
	"fmt"
)

type Order struct {
	Id    int
	Total float64
}

func (o Order) CreateOrder(ctx context.Context) error {

	fmt.Printf("Order Instance - Create order executed \n")
	return nil
}

func (o Order) RollbackCreateOrder(ctx context.Context) error {

	fmt.Printf("Order Instance - Create order rolled back \n")
	return nil
}
