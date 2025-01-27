package models

import (
	"sync"
	"time"
)

type CachedValue[T comparable] struct {
	value       T
	latestError error
	cachedDate  time.Time
	validity    time.Duration
	updaterFn   func(self *CachedValue[T])
	mu          *sync.Mutex
}

func NewCachedValue[T comparable](
	value T,
	err error,
	cachedDate time.Time,
	validity time.Duration,
	updaterFn func(*CachedValue[T]),
) *CachedValue[T] {
	return &CachedValue[T]{
		value:       value,
		latestError: err,
		cachedDate:  cachedDate,
		validity:    validity,
		updaterFn:   updaterFn,
		mu:          &sync.Mutex{},
	}
}

func (c *CachedValue[T]) Get() (T, error) {
	shouldUpdate := false
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		if shouldUpdate {
			c.updaterFn(c)
		}
	}()
	shouldUpdate = c.cachedDate.Add(c.validity).Before(time.Now())
	return c.value, c.latestError
}

func (c *CachedValue[T]) Set(value T, err error) bool {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.cachedDate = time.Now()
	updated := c.value == value
	c.value = value
	c.latestError = err

	return updated
}
