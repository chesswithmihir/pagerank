package graph

import (
	"sync"
)

// Mapper maps string URLs to uint64 IDs and vice-versa.
type Mapper struct {
	mu       sync.RWMutex
	strToID  map[string]uint64
	idToStr  []string
	nextID   uint64
}

func NewMapper() *Mapper {
	return &Mapper{
		strToID: make(map[string]uint64),
		idToStr: make([]string, 0),
		nextID:  0,
	}
}

// GetID returns the ID for a URL, creating it if it doesn't exist.
func (m *Mapper) GetID(url string) uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()

	if id, exists := m.strToID[url]; exists {
		return id
	}

	id := m.nextID
	m.strToID[url] = id
	m.idToStr = append(m.idToStr, url)
	m.nextID++
	return id
}

// GetURL returns the URL for a given ID.
func (m *Mapper) GetURL(id uint64) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if id >= uint64(len(m.idToStr)) {
		return ""
	}
	return m.idToStr[id]
}

func (m *Mapper) Size() uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nextID
}
