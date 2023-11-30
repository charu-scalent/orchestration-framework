package main

import (
	"context"
	"errors"
	"fmt"
)

type Inventory struct {
	Id    int
	Total int
}

func (i Inventory) UpdateInventory(ctx context.Context) (interface{}, error) {

	fmt.Printf("Inventory Instance - Update Inventory executed \n")
	return "", errors.New("insufficient balance")
}

func (i Inventory) RollbackUpdateInventory(ctx context.Context) error {
	fmt.Printf("Inventory Instance - Update Inventory rolled back \n")
	return nil
}
