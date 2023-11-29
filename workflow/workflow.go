package workflow

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

const TOTAL_EXPECTED_RESPONSES = 2

type Workflow struct {
	Steps          []Step
	IdempotentKey  string
	IdempotentInst Idempotent
}

type Step struct {
	Instance        interface{}
	Method          string
	RollbackMethod  string
	IsMandatoryStep bool
	IsExecuted      bool
	StepResult      interface{}
	StepError       error
}

func CreateWorkflow(idempotentKey string, idempotentInst Idempotent) *Workflow {
	return &Workflow{
		IdempotentKey:  idempotentKey,
		IdempotentInst: idempotentInst,
	}
}

func (w *Workflow) Register(instance interface{}, method, rollbackMethod string, isMandatoryStep bool) {

	step := Step{Instance: instance,
		Method:          method,
		RollbackMethod:  rollbackMethod,
		IsMandatoryStep: isMandatoryStep,
	}
	w.Steps = append(w.Steps, step)
}

func (w *Workflow) Execute(ctx context.Context) error {
	w.IdempotentInst.Save(w.IdempotentKey, w.Steps)
	for _, step := range w.Steps {
		err := w.executeStep(ctx, w.IdempotentKey, step)
		if err != nil && step.IsMandatoryStep {
			//TODO: instantiate rollback procedure
			return err
		}
	}
	return nil
}

func (w *Workflow) executeStep(ctx context.Context, idempotentKey string, step Step) error {

	if idempotentKey == "" {
		return errors.New("missing Idempotent-Key")
	}

	if w.IdempotentInst.IsStepAlreadyExecuted(ctx, step.Method, idempotentKey) {
		fmt.Printf("Step %s skipped as it has already been executed with idempotent key: %s\n", step.Method, idempotentKey)
		return nil
	}

	var arg reflect.Value = reflect.ValueOf(ctx)
	var ref []reflect.Value
	ref = append(ref, arg)

	method := reflect.ValueOf(step.Instance).MethodByName(step.Method)
	response := method.Call(ref)

	if len(response) != TOTAL_EXPECTED_RESPONSES {
		return fmt.Errorf("method %s did not return 2 values (result, error)", step.Method)
	}

	var err error
	var ok bool

	responseError := response[1].Interface()
	if responseError != nil {
		if err, ok = responseError.(error); !ok {
			return fmt.Errorf("error occured while executing the method: %s, error: %w", step.Method, err)
		}
	}

	step.StepResult = response[0]

	w.IdempotentInst.MarkStepAsExecuted(ctx, idempotentKey, step.Method, step.StepResult, step.StepError)

	return err
}
