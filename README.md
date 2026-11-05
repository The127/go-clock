# go-clock

go-clock gives Go code the current time through an interface, so a test can set the time instead of waiting for it.

Code that calls `time.Now` directly is hard to test. A test can't check what happens when a token expires without sleeping until it does. With go-clock, the code asks a `clock.Service` for the time. In production that's the system clock, and in a test it's a mock clock that you set.

## Install

```sh
go get github.com/The127/go-clock
```

## Use it

Pass the clock to the code that needs the time:

```go
type Tokens struct {
	clock clock.Service
}

func NewTokens(c clock.Service) *Tokens {
	return &Tokens{clock: c}
}

func (t *Tokens) Expired(expiry time.Time) bool {
	return t.clock.Now().After(expiry)
}
```

In production code, give it the system clock:

```go
tokens := NewTokens(clock.NewSystemClock())
```

In a test, give it a mock clock and move its time forward:

```go
func TestTokensExpire(t *testing.T) {
	start := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	mock, setTime := clock.NewMockClock(start)
	tokens := NewTokens(mock)

	setTime(start.Add(2 * time.Hour))

	if !tokens.Expired(start.Add(time.Hour)) {
		t.Error("the token should have expired")
	}
}
```

`NewMockClock` returns the clock and a function that sets its time. You can call both from several goroutines at once.

go-clock has no dependencies outside Go's standard library.

## Contribute

[CONTRIBUTING.md](CONTRIBUTING.md) explains how to set up the tools, how to sign off your commits and how releases work.

## License

go-clock is licensed under the [MIT License](LICENSE.md).
