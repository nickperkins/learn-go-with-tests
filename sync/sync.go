package sync

import "sync"

// A Counter is used for counting
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a reference to a new Counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc will increment the counter by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current Counter value
func (c *Counter) Value() int {
	return c.value
}
