package lib

import "sync"

// SafeColumnData is a thread-safe structure for storing column data.
type SafeColumnData struct {
	data map[string][]string
	mu   sync.RWMutex
}

// NewSafeColumnData initializes a new SafeColumnData instance.
func NewSafeColumnData(headers []string) *SafeColumnData {
	data := make(map[string][]string)
	for _, header := range headers {
		data[header] = make([]string, 0)
	}
	return &SafeColumnData{
		data: data,
	}
}

// Add appends a value to the specified column in a thread-safe manner.
func (s *SafeColumnData) Add(header, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[header] = append(s.data[header], value)
}

// GetData returns the underlying data map.
func (s *SafeColumnData) GetData() map[string][]string {
	return s.data
}
