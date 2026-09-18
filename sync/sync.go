package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// c.value++ is not safe when multiple goroutines access the same Counter concurrently. Goroutine A and B could read the same value of c.value, increment it, and write it back, resulting in a lost update. A and B must be different count. To fix this, we need to use a mutex to ensure that only one goroutine can access the critical section of code at a time.

	// What this means is any goroutine calling Inc will acquire the lock on Counter if they are first. All the other goroutines will have to wait for it to be Unlocked before getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
