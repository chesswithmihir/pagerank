package utils

const (
	// DampingFactor is the probability that the random surfer follows a link.
	// 1 - DampingFactor is the probability they teleport to a random page.
	DampingFactor = 0.85

	// Epsilon is the convergence threshold. If the change in rank is less than this, we stop.
	Epsilon = 1e-9

	// MaxIterations prevents infinite loops if it doesn't converge.
	MaxIterations = 100
)
