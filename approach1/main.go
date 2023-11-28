package main

import (
	"context"
	"fmt"

	"github.com/scalent-io/orchestration-framework/approach1/workflow"
)

func main() {

	w := workflow.CreateWorkflow()
	var order Order

	order.Id = 1
	order.Total = 100

	w.Register(order, "CreateOrder", "RollbackCreateOrder", true)
	w.Register(order, "DeductBalance", "RollbackDeductBalance", true)

	w.Execute(context.Background(), "Idempotent-Key")
	fmt.Println("Main executed")
}
