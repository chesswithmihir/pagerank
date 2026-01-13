package math

import (
	"math"
)

// ======================================================================================
// LEARNING TASK #3: Vector Operations
// ======================================================================================

// L1Difference calculates the sum of absolute differences between two vectors.
// This is used to check for convergence.
// Formula: sum(|v1[i] - v2[i]|) for all i
func L1Difference(v1, v2 []float64) float64 {
	if len(v1) != len(v2) {
		panic("vector lengths do not match")
	}

	var diff float64 = 0.0

	// TODO: Implement the L1 Difference loop.
	// 1. Iterate through the vectors.
	// 2. Add math.Abs(v1[i] - v2[i]) to diff.
	
	// Silence unused import error for boilerplate
	_ = math.Abs(0)

	// --- YOUR CODE HERE ---
	for i, v1_val := range v1 {
		v2_val := v2[i]
		diff += math.Abs(v1_val - v2_val)
	}

	return diff
}

// InitializeUniform creates a vector of size N where every element is 1/N.
func InitializeUniform(size int) []float64 {
	v := make([]float64, size)
	val := 1.0 / float64(size)
	for i := range v {
		v[i] = val
	}
	return v
}
