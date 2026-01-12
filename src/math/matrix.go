package math

// CSRMatrix represents a sparse graph in Compressed Sparse Row format.
// NOTE: For our PageRank engine, this effectively stores the "Incoming Links" (Reverse Graph).
// Row 'i' corresponds to node 'i', and the columns listed in row 'i' are the nodes that point TO 'i'.
type CSRMatrix struct {
	// Values holds the weights (probabilities).
	// In standard PageRank, if node J links to I, the weight is 1.0 / OutDegree(J).
	Values []float64

	// ColumnIndices holds the column index (Source Node ID) for each value.
	ColumnIndices []uint64

	// RowPtrs indicates where each row starts in the Values/ColumnIndices arrays.
	// Row i's data is at indices [ RowPtrs[i], RowPtrs[i+1] ).
	// Length is NumNodes + 1.
	RowPtrs []uint64

	Rows int // Number of nodes
}

// NewCSRMatrix creates an empty matrix structure.
func NewCSRMatrix(rows int, values []float64, cols []uint64, rowPtrs []uint64) *CSRMatrix {
	return &CSRMatrix{
		Values:        values,
		ColumnIndices: cols,
		RowPtrs:       rowPtrs,
		Rows:          rows,
	}
}

// ======================================================================================
// LEARNING TASK #1: Sparse Matrix-Vector Multiplication
// ======================================================================================

// Multiply performs the sparse matrix-vector multiplication operation: result = M * v.
// It is the core of the "Pull" approach to PageRank.
//
// For each node 'i' in the graph (each row in our matrix), this function calculates
// the total rank flowing into it from all the nodes that have outbound links to 'i'.
//
//   - m (*CSRMatrix): The receiver. This is the graph itself, stored in CSR format
//     where rows represent target nodes and column indices represent source nodes.
//   - v ([]float64): A vector representing the current PageRank of every node in the graph.
//
// Returns a new vector ([]float64) that holds the next PageRank values before applying
// the damping factor.
func (m *CSRMatrix) Multiply(v []float64) []float64 {
	result := make([]float64, m.Rows)

	// TODO: Implement the Sparse Matrix-Vector Multiplication.
	//
	// Guide:
	// 1. Iterate over every row 'i' from 0 to m.Rows-1.
	// 2. For each row, determine the start and end indices in the Values/ColumnIndices arrays
	//    using m.RowPtrs[i] and m.RowPtrs[i+1].
	// 3. Iterate from start to end:
	//    a. Get the 'sourceNode' (m.ColumnIndices[k])
	//    b. Get the 'weight' (m.Values[k])
	//    c. Add (weight * v[sourceNode]) to result[i].
	
	// --- YOUR CODE HERE ---
	
	return result
}
