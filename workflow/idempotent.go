package workflow

import "context"

type Idempotent interface {
	Save(idempotentKey string, steps []Step)
	IsStepAlreadyExecuted(ctx context.Context, step, idempotentKey string) bool
	MarkStepAsExecuted(ctx context.Context, idempotentKey, step string, result interface{}, err error)
}
