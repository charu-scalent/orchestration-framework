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

	fmt.Printf("Create order executed \n")
	return nil
}

func (o Order) RollbackCreateOrder(ctx context.Context) error {

	fmt.Printf("Create order rolled back \n")
	return nil
}

func (o Order) DeductBalance(ctx context.Context) error {

	fmt.Printf("Deduct balance executed \n")
	return nil
}

func (o Order) RollbackDeductBalance(ctx context.Context) error {

	fmt.Printf("Deduct balance rolled back \n")
	return nil
}
