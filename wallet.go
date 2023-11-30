package main

import (
	"context"
	"errors"
	"fmt"
)

type WalletInterface interface {
	DeductBalance(ctx context.Context) error
	RollbackDeductBalance(ctx context.Context) error
}

type Wallet struct {
	Id      int
	Balance float64
}

func (w Wallet) DeductBalance(ctx context.Context) (interface{}, error) {
	fmt.Printf("Wallet Instance - Deduct balance executed \n")
	return "", errors.New("Wallet deduction failed")
}

func (w Wallet) RollbackDeductBalance(ctx context.Context) error {
	fmt.Printf("Wallet Instance - Deduct balance rolled back \n")
	return nil
}
