package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/scalent-io/ecommerce/workflow"
)

const idmp_key_prefix = "IDMP:"

type IdempotentOp struct {
	redisClient *redis.Client
}

func (o IdempotentOp) Save(idempotentKey string, steps []workflow.Step) {
	idmp_key := getIdempotentRedisKey(idempotentKey)
	dataMap := make(map[string]workflow.Step)

	for _, step := range steps {
		dataMap[step.Method] = step
	}

	dataInstance, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println(err.Error())
	}

	redisResult := o.redisClient.Set(idmp_key, dataInstance, time.Duration(time.Hour*24))
	if redisResult.Err() != nil {
		fmt.Println(redisResult.Err().Error())
	}

	fmt.Println("IdempotentOp Save called, steps: ", steps)
}

func (o IdempotentOp) IsStepAlreadyExecuted(ctx context.Context, stepMethod, idempotentKey string) bool {
	dataMap := make(map[string]workflow.Step)
	idmp_key := getIdempotentRedisKey(idempotentKey)

	redisResult, err := o.redisClient.Get(idmp_key).Result()
	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(redisResult), &dataMap)
	if err != nil {
		return false
	}

	if step, ok := dataMap[stepMethod]; ok {
		return step.IsExecuted
	}

	fmt.Printf("%+v", dataMap)

	return false
}

func (idmp IdempotentOp) MarkStepAsExecuted(ctx context.Context, idempotentKey, stepMethod string, result interface{}, stepError error) {

	dataMap := make(map[string]workflow.Step)
	idmp_key := getIdempotentRedisKey(idempotentKey)

	redisResult, err := idmp.redisClient.Get(idmp_key).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(redisResult), &dataMap)
	if err != nil {
		fmt.Println(err.Error())
	}

	if step, ok := dataMap[stepMethod]; ok {
		step.IsExecuted = true
		step.StepResult = result
		step.StepError = stepError
		dataMap[stepMethod] = step
	}

	dataInstance, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println(err.Error())
	}

	idmp.redisClient.Set(idmp_key, dataInstance, time.Duration(time.Hour*24))
}

func getIdempotentRedisKey(idempotentKey string) string {
	//prefix to key idmp_key and use this key to store in redis
	return idmp_key_prefix + idempotentKey
}

// func
