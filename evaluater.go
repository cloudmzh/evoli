package darwin

// Evaluater computes and set individual Fitness
type Evaluater interface {
	Evaluate(Individual) (Fitness float32)
}