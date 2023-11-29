package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/scalent-io/orchestration-framework/approach1/workflow"
)

func main() {

	idempotentKey := uuid.New().String()
	var idempotentOp IdempotentOp

	w := workflow.CreateWorkflow(idempotentKey, idempotentOp)
	order := Order{Id: 1, Total: 120}
	wallet := Wallet{Id: 4, Balance: 200}

	w.Register(order, "CreateOrder", "RollbackCreateOrder", true)
	w.Register(wallet, "DeductBalance", "RollbackDeductBalance", true)

	w.Execute(context.Background())
	fmt.Println("Main executed")
}
