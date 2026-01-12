package graph

import (
	"project-eigenweb/src/math"
)

// Loader handles reading graph data from disk.
type Loader struct {
	// TODO: Add file path or configuration
}

// LoadFromEdgeList reads a file where each line is "SourceURL DestURL".
// It returns the CSRMatrix (Transpose Graph) and the Mapper.
func LoadFromEdgeList(filePath string) (*math.CSRMatrix, *Mapper, error) {
	// Placeholder for future implementation.
	// 1. Read file line by line.
	// 2. Use Mapper to get IDs.
	// 3. Build Adjacency List (In-links).
	// 4. Convert to CSR.
	return nil, nil, nil
}
