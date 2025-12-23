# go-clock

A simple and elegant Go library for time abstraction, designed to make testing time-dependent code easier.

## Overview

`go-clock` provides a clean interface for working with time in Go applications. It includes both a production implementation using the system clock and a mock implementation for testing purposes.

This library is considered **complete**, but any additions are welcome.

## Features

- **Simple Interface**: A clean `Service` interface with a `Now()` method
- **System Clock**: Production-ready implementation using `time.Now()`
- **Mock Clock**: Controllable time implementation for testing
- **Zero Dependencies**: Pure Go with only the standard library
- **MIT Licensed**: Free to use in any project

## Installation

```bash
go get github.com/The127/go-clock
```

## Usage

### Using the System Clock

For production code, use the system clock implementation:

```go
package main

import (
    "fmt"
    "github.com/The127/go-clock"
)

func main() {
    clockService := clock.NewSystemClock()
    currentTime := clockService.Now()
    fmt.Println("Current time:", currentTime)
}
```

### Using the Mock Clock for Testing

For testing, use the mock clock to control time:

```go
package main

import (
    "testing"
    "time"
    "github.com/The127/go-clock"
)

func TestTimeDependent(t *testing.T) {
    // Create a mock clock with a specific time
    fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
    clockService, setTime := clock.NewMockClock(fixedTime)
    
    // Use the clock in your code
    if clockService.Now() != fixedTime {
        t.Error("Expected fixed time")
    }
    
    // Advance time as needed
    newTime := fixedTime.Add(1 * time.Hour)
    setTime(newTime)
    
    if clockService.Now() != newTime {
        t.Error("Expected new time")
    }
}
```

### Dependency Injection Pattern

The recommended pattern is to inject the clock service into your components:

```go
type MyService struct {
    clock clock.Service
}

func NewMyService(clockService clock.Service) *MyService {
    return &MyService{
        clock: clockService,
    }
}

func (s *MyService) DoSomethingWithTime() {
    currentTime := s.clock.Now()
    // Use currentTime...
}
```

This allows you to use the system clock in production and the mock clock in tests.

## API Reference

### `Service` Interface

```go
type Service interface {
    Now() time.Time
}
```

The main interface for time retrieval.

### `NewSystemClock() Service`

Creates a new system clock that returns the current time using `time.Now()`.

### `NewMockClock(now time.Time) (Service, TimeSetterFn)`

Creates a new mock clock initialized to the specified time. Returns both the clock service and a function to update the time.

### `TimeSetterFn`

```go
type TimeSetterFn func(time.Time)
```

A function type used to update the time in a mock clock.

## Contributing

This library is considered complete for its intended purpose, but contributions are welcome! If you have ideas for improvements or additional features, feel free to:

- Open an issue to discuss your ideas
- Submit a pull request with your changes
- Report any bugs or issues you encounter

## License

This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md) for details.

## Author

The127
