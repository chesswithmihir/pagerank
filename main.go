package main

import (
	"fmt"
	"project-eigenweb/src/engine"
	"project-eigenweb/src/math"
)

func main() {
	fmt.Println("Welcome to Project EigenWeb!")

	// 1. Create a Tiny Graph (The "Tiny Web Test")
	// 0 -> 1
	// 1 -> 2
	// 2 -> 0
	// This is a simple loop. Ranks should be equal (0.333...). 
	//
	// In Reverse Graph (Source -> Target becomes WhoPointsTo -> Node):
	// Node 0 is pointed to by 2.
	// Node 1 is pointed to by 0.
	// Node 2 is pointed to by 1.
	
	// Values (weights): All 1.0 (since out-degree is 1 for all)
	values := []float64{1.0, 1.0, 1.0}
	
	// ColumnIndices (The sources): 
	// Row 0's source: 2
	// Row 1's source: 0
	// Row 2's source: 1
	cols := []uint64{2, 0, 1}
	
	// RowPtrs:
	// Row 0 starts at 0
	// Row 1 starts at 1
	// Row 2 starts at 2
	// End is 3
	rowPtrs := []uint64{0, 1, 2, 3}
	
graph := math.NewCSRMatrix(3, values, cols, rowPtrs)

	// 2. Initialize Engine
	eng := engine.NewEngine(graph)

	// 3. Run
	eng.Run()

	// 4. Print Results
	for i, rank := range eng.CurrentRanks {
		fmt.Printf("Node %d Rank: %.4f\n", i, rank)
	}
}
