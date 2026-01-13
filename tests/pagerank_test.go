package tests

import (
	"math"
	"testing"

	projectMath "project-eigenweb/src/math"
	"project-eigenweb/src/engine"
)

const floatTolerance = 1e-9

// ======================================================================================
// TEST TASK #3: Vector Operations (L1Difference)
// ======================================================================================

func TestL1Difference(t *testing.T) {
	v1 := []float64{0.1, 0.5, 0.4}
	v2 := []float64{0.2, 0.4, 0.4}
	
	// |0.1 - 0.2| + |0.5 - 0.4| + |0.4 - 0.4|
	// 0.1 + 0.1 + 0 = 0.2
	expected := 0.2

	result := projectMath.L1Difference(v1, v2)

	if math.Abs(result-expected) > floatTolerance {
		t.Errorf("L1Difference incorrect. Expected %.5f, got %.5f", expected, result)
	}
}

func TestL1Difference_Zero(t *testing.T) {
	v1 := []float64{0.3, 0.3}
	v2 := []float64{0.3, 0.3}
	expected := 0.0
	
	result := projectMath.L1Difference(v1, v2)
	if math.Abs(result-expected) > floatTolerance {
		t.Errorf("L1Difference incorrect for identical vectors. Expected %.5f, got %.5f", expected, result)
	}
}

// ======================================================================================
// TEST TASK #1: Sparse Matrix-Vector Multiplication
// ======================================================================================

func TestCSRMultiply(t *testing.T) {
	// Let's build a simple 3-node graph (transpose/incoming links representation).
	// Node 0 receives from Node 1 (weight 0.5) and Node 2 (weight 1.0)
	// Node 1 receives from Node 0 (weight 1.0)
	// Node 2 receives from Node 1 (weight 0.5)
	
	// This corresponds to a standard adjacency graph:
	// 0 -> 1 (100% of 0's output)
	// 1 -> 0 (50% of 1's output)
	// 1 -> 2 (50% of 1's output)
	// 2 -> 0 (100% of 2's output)

	// Matrix M (transpose) would be:
	//    0    1    2
	// 0 [0.0, 0.5, 1.0]  <-- Node 0 gets 0.5 from 1, 1.0 from 2
	// 1 [1.0, 0.0, 0.0]  <-- Node 1 gets 1.0 from 0
	// 2 [0.0, 0.5, 0.0]  <-- Node 2 gets 0.5 from 1

	values := []float64{0.5, 1.0, 1.0, 0.5}
	colIndices := []uint64{1, 2, 0, 1}
	
	// Row 0 has 2 entries (indices 0, 1)
	// Row 1 has 1 entry (index 2)
	// Row 2 has 1 entry (index 3)
	rowPtrs := []uint64{0, 2, 3, 4}

	csr := projectMath.NewCSRMatrix(3, values, colIndices, rowPtrs)

	// Vector v (current ranks)
	// Let's say everyone starts with 1.0 for easy math
	v := []float64{1.0, 1.0, 1.0}

	// Expected Result: M * v
	// Row 0: 0.5*1 + 1.0*1 = 1.5
	// Row 1: 1.0*1 = 1.0
	// Row 2: 0.5*1 = 0.5
	expected := []float64{1.5, 1.0, 0.5}

	result := csr.Multiply(v)

	if len(result) != 3 {
		t.Fatalf("Result vector length wrong. Expected 3, got %d", len(result))
	}

	for i := range expected {
		if math.Abs(result[i]-expected[i]) > floatTolerance {
			t.Errorf("Multiply row %d incorrect. Expected %.5f, got %.5f", i, expected[i], result[i])
		}
	}
}

// ======================================================================================
// TEST TASK #2: The PageRank Step
// ======================================================================================

func TestPageRankStep(t *testing.T) {
	// Re-use the graph from above.
	values := []float64{0.5, 1.0, 1.0, 0.5}
	colIndices := []uint64{1, 2, 0, 1}
	rowPtrs := []uint64{0, 2, 3, 4}
	csr := projectMath.NewCSRMatrix(3, values, colIndices, rowPtrs)

	engineObj := engine.NewEngine(csr)
	
	// Override default config for predictable testing
	engineObj.DampingFactor = 0.85
	// Force current ranks to be uniform (0.333...) which NewEngine does, but let's be explicit if we want specific values
	// Let's set specific values to test calculation easier.
	engineObj.CurrentRanks = []float64{1.0, 1.0, 1.0} 
	// NOTE: Ranks usually sum to 1, but the formula works for any magnitude if we don't normalize.
	// For this test, we assume unnormalized for simplicity of manual calc.
	
	// Incoming Ranks (calculated in Multiply test): [1.5, 1.0, 0.5]
	
	// Formula: PR(i) = (1-d)/N + d * (Incoming + DanglingWeight/N)
	// d = 0.85
	// N = 3
	// (1-d)/N = 0.15 / 3 = 0.05
	
	// We assume DanglingWeight = 0 for this basic test (no dead ends in our graph).
	
	// Expected Next Ranks:
	// Node 0: 0.05 + 0.85 * (1.5) = 0.05 + 1.275 = 1.325
	// Node 1: 0.05 + 0.85 * (1.0) = 0.05 + 0.85  = 0.900
	// Node 2: 0.05 + 0.85 * (0.5) = 0.05 + 0.425 = 0.475

	nextRanks, _ := engineObj.Step()

	expected := []float64{1.325, 0.900, 0.475}

	for i := range expected {
		if math.Abs(nextRanks[i]-expected[i]) > floatTolerance {
			t.Errorf("Step node %d incorrect. Expected %.5f, got %.5f", i, expected[i], nextRanks[i])
		}
	}
}
