package workflow

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Workflow struct {
	steps          []Step
	idempotentKey  string
	idempotentInst Idempotent
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
		idempotentKey:  idempotentKey,
		idempotentInst: idempotentInst,
	}
}

func (w *Workflow) Register(instance interface{}, method, rollbackMethod string, isMandatoryStep bool) {

	step := Step{Instance: instance,
		Method:          method,
		RollbackMethod:  rollbackMethod,
		IsMandatoryStep: isMandatoryStep,
	}
	w.steps = append(w.steps, step)
}

func (w *Workflow) Execute(ctx context.Context) error {
	w.idempotentInst.Save(w.idempotentKey, w.steps)
	for _, step := range w.steps {
		err := w.executeStep(ctx, w.idempotentKey, step)
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

	if w.idempotentInst.IsStepAlreadyExecuted(ctx, step.Method, idempotentKey) {
		fmt.Printf("Step %s skipped as it has already been executed with idempotent key: %s\n", step.Method, idempotentKey)
		return nil
	}

	var arg reflect.Value = reflect.ValueOf(ctx)
	var ref []reflect.Value
	ref = append(ref, arg)

	method := reflect.ValueOf(step.Instance).MethodByName(step.Method)
	response := method.Call(ref) //TODO: handle error and start rolling back if it's a mandatory step

	if len(response) != 2 {
		return errors.New(fmt.Sprintf("Method %s did not return 2 values (result, error)", step.Method))
	}

	errValue := response[1].Interface()
	var err error

	if errValue != nil {
		var ok bool
		if err, ok = errValue.(error); !ok {
			return errors.New(fmt.Sprintf("Method %s did not return an error as the second value", step.Method))
		}
	}

	step.StepResult = response[0]

	w.idempotentInst.MarkStepAsExecuted(ctx, idempotentKey, step.Method, step.StepResult, step.StepError)

	return err
}
