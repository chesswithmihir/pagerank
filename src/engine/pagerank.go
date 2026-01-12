package engine

import (
	"fmt"
	"project-eigenweb/src/math"
	"project-eigenweb/src/utils"
)

type PageRankEngine struct {
	Graph          *math.CSRMatrix
	DampingFactor  float64
	CurrentRanks   []float64
	DanglingWeight float64 // Total rank of dangling nodes (nodes with no out-links) to redistribute
}

func NewEngine(graph *math.CSRMatrix) *PageRankEngine {
	n := graph.Rows
	return &PageRankEngine{
		Graph:         graph,
		DampingFactor: utils.DampingFactor,
		CurrentRanks:  math.InitializeUniform(n),
	}
}

// ======================================================================================
// LEARNING TASK #2: The PageRank Step
// ======================================================================================

// Step performs one iteration of the Power Method.
// Returns the next rank vector and the delta (convergence error).
func (e *PageRankEngine) Step() ([]float64, float64) {
	numNodes := float64(e.Graph.Rows)
	
	// 1. Calculate the rank flow from links: M * v_current
	// Use your math.Multiply implementation!
	incomingRanks := e.Graph.Multiply(e.CurrentRanks)

	// 2. Prepare the next rank vector
	nextRanks := make([]float64, len(incomingRanks))

	// TODO: Implement the core PageRank formula application.
	//
	// Formula for each node i:
	// PR(i) = (1 - d) / N  +  d * ( IncomingRank(i) + DanglingWeight / N )
	//
	// Where:
	// d = e.DampingFactor
	// N = numNodes
	// IncomingRank(i) = incomingRanks[i] (calculated above)
	//
	// Note: We are simplifying DanglingWeight for this exercise. 
	// Assume 'e.DanglingWeight' is already calculated or just use 0.0 if you want to keep it simple first.
	// (Real implementation requires calculating total rank lost to dangling nodes in previous step).
	
	auxiliaryConstant := (1.0 - e.DampingFactor) / numNodes
	
	// Iterate over nextRanks and apply the formula
	for i := range nextRanks {
		// --- YOUR CODE HERE ---
		// nextRanks[i] = ...
		_ = i // remove this when you code
		_ = auxiliaryConstant 
	}

	// 3. Calculate Convergence Delta (L1 Norm)
	// Use your math.L1Difference implementation!
	delta := math.L1Difference(nextRanks, e.CurrentRanks)

	return nextRanks, delta
}

func (e *PageRankEngine) Run() {
	fmt.Println("Starting PageRank...")
	
	for i := 0; i < utils.MaxIterations; i++ {
		next, delta := e.Step()
		
		fmt.Printf("Iteration %d: Delta = %.10f\n", i+1, delta)
		
e.CurrentRanks = next
		if delta < utils.Epsilon {
			fmt.Println("Converged!")
			break
		}
	}
}
