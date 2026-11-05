// Package clock tells the time through an interface, so tests can set it.
package clock

import (
	"sync"
	"time"
)

// Service tells the current time.
type Service interface {
	Now() time.Time
}

// TimeSetterFn sets the time that a mock clock returns.
type TimeSetterFn func(time.Time)

type mockClock struct {
	mu  sync.Mutex
	now time.Time
}

func (m *mockClock) Now() time.Time {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.now
}

func (m *mockClock) set(now time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.now = now
}

// NewMockClock returns a clock that always returns now, and a function that
// changes the time it returns. Both are safe to call from several goroutines.
func NewMockClock(now time.Time) (Service, TimeSetterFn) {
	m := &mockClock{now: now}

	return m, m.set
}

type systemClock struct{}

// NewSystemClock returns a clock that returns time.Now.
func NewSystemClock() Service {
	return systemClock{}
}

func (systemClock) Now() time.Time {
	return time.Now()
}
