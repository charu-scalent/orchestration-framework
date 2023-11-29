package main

import (
	"context"
	"fmt"

	"github.com/scalent-io/orchestration-framework/approach1/workflow"
)

const idmp_key_prefix = "IDMP_KEY:"

type IdempotentOp struct {
}

func (o IdempotentOp) Save(idempotentKey string, steps []workflow.Step) {
	// getIdempotentRedisKey(idempotentKey string)
	//
	fmt.Println("IdempotentOp Save called, steps: ", steps)
}

func (o IdempotentOp) IsStepAlreadyExecuted(ctx context.Context, step, idempotentKey string) bool {
	return false
}

func (o IdempotentOp) MarkStepAsExecuted(ctx context.Context, step, idempotentKey string) {

}

func getIdempotentRedisKey(idempotentKey string) string {
	//prefix to key idmp_key and use this key to store in redis
	return ""
}
