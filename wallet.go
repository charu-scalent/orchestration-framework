package main

import (
	"context"
	"fmt"
)

type Wallet struct {
	Id      int
	Balance float64
}

func (w Wallet) DeductBalance(ctx context.Context) error {

	fmt.Printf("Wallet Instance - Deduct balance executed \n")
	return nil
}

func (w Wallet) RollbackDeductBalance(ctx context.Context) error {
	fmt.Printf("Wallet Instance - Deduct balance rolled back \n")
	return nil
}
