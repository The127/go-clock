package clock

import "time"

// Service defines an interface for retrieving the current time.
type Service interface {
	Now() time.Time
}

// TimeSetterFn defines a function type used to update a time value in a specific context or object.
type TimeSetterFn func(time.Time)

// mockClock is a type that implements the Service interface for mocking time-related functionality.
type mockClock struct {
	now time.Time
}

// Now returns the current mocked time stored in the mockClock instance.
func (m *mockClock) Now() time.Time {
	return m.now
}

// NewMockClock creates a mock Service with a specified initial time and returns it along with a function to update the time.
func NewMockClock(now time.Time) (Service, TimeSetterFn) {
	service := mockClock{
		now: now,
	}

	return &service, func(t time.Time) {
		service.now = t
	}
}

// clockService is a private implementation of the Service interface that provides the current time.
type systemClock struct{}

// NewSystemClock creates and returns a new instance of a Service implementation for handling time-related operations with the system clock using time.Now().
func NewSystemClock() Service {
	return &systemClock{}
}

// Now returns the current local time using the system clock.
func (c *systemClock) Now() time.Time {
	return time.Now()
}
