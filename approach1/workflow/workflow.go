package workflow

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Workflow struct {
	steps []Step
}

func CreateWorkflow() *Workflow {
	return &Workflow{}
}

type Step struct {
	Instance        interface{}
	Method          string
	RollbackMethod  string
	IsMandatoryStep bool
}

func (w *Workflow) Register(instance interface{}, method, rollbackMethod string, isMandatoryStep bool) {

	step := Step{Instance: instance,
		Method:          method,
		RollbackMethod:  rollbackMethod,
		IsMandatoryStep: isMandatoryStep,
	}
	w.steps = append(w.steps, step)
}

func (w *Workflow) Execute(ctx context.Context, idempotentKey string) error {
	for _, step := range w.steps {
		if err := w.executeStep(ctx, idempotentKey, step); err != nil {
			return err
		}
	}
	return nil
}

func (w *Workflow) executeStep(ctx context.Context, idempotentKey string, step Step) error {

	if idempotentKey == "" {
		return errors.New("missing Idempotent-Key")
	}

	if isStepAlreadyExecuted(ctx, step.Method, idempotentKey) {
		fmt.Printf("Step %s skipped as it has already been executed with idempotent key: %s\n", step.Method, idempotentKey)
		return nil
	}

	var b reflect.Value = reflect.ValueOf(ctx)
	var ref []reflect.Value
	ref = append(ref, b)

	method := reflect.ValueOf(step.Instance).MethodByName(step.Method)
	method.Call(ref) //TODO: handle error and start rolling back if it's a mandatory step

	markStepAsExecuted(ctx, step.Method, idempotentKey)

	return nil
}

func isStepAlreadyExecuted(ctx context.Context, step, idempotentKey string) bool {

	return false
}

func markStepAsExecuted(ctx context.Context, step, idempotentKey string) {

}
