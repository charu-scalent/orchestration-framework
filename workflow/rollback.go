package workflow

import "context"

func (w *Workflow) rollBack(ctx context.Context, idempotentKey string) error {
	for _, step := range w.steps {
		if err := w.rollbackStep(ctx, idempotentKey, step); err != nil {
			return err
		}
	}
	return nil
}

func (w *Workflow) rollbackStep(ctx context.Context, idempotentKey string, step Step) error {

	//TODO: rollback step to be handled here
	return nil
}
