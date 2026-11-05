package clock_test

import (
	"sync"
	"testing"
	"time"

	"github.com/The127/go-clock"
)

var start = time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

func TestMockClockReturnsTheTimeItWasCreatedWith(t *testing.T) {
	t.Parallel()

	// arrange
	service, _ := clock.NewMockClock(start)

	// act
	now := service.Now()

	// assert
	if !now.Equal(start) {
		t.Errorf("Now() = %v, want %v", now, start)
	}
}

func TestMockClockReturnsTheTimeTheSetterSets(t *testing.T) {
	t.Parallel()

	// arrange
	service, setTime := clock.NewMockClock(start)
	later := start.Add(time.Hour)

	// act
	setTime(later)

	// assert
	if now := service.Now(); !now.Equal(later) {
		t.Errorf("Now() = %v, want %v", now, later)
	}
}

func TestMockClockCanBeReadWhileItsTimeIsSet(t *testing.T) {
	t.Parallel()

	// arrange
	service, setTime := clock.NewMockClock(start)

	var wg sync.WaitGroup

	// act
	wg.Go(func() { setTime(start.Add(time.Hour)) })
	wg.Go(func() { _ = service.Now() })
	wg.Wait()

	// assert
	// The race detector fails the test if Now and the setter race.
}
