package workflow

type Idempotent interface {
	GenerateNewKey()
}

type IdempotentStep interface {
	GetSteps(key string)
	UpdateSteps(key, step, status string)
}
