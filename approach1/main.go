package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/scalent-io/orchestration-framework/approach1/workflow"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	idempotentKey := uuid.New().String()
	idempotentOp := IdempotentOp{
		redisClient: redisClient,
	}

	w := workflow.CreateWorkflow(idempotentKey, idempotentOp)
	order := Order{Id: 1, Total: 120}
	wallet := Wallet{Id: 4, Balance: 200}

	w.Register(order, "CreateOrder", "RollbackCreateOrder", true)
	w.Register(wallet, "DeductBalance", "RollbackDeductBalance", true)

	w.Execute(context.Background())
	fmt.Println("Main executed")
}
