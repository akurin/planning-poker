package domain

import (
	"time"
)

type fakeClock struct {
	now time.Time
}

func NewFakeClock(opts ...FakeClockOption) Clock {
	defaultNow := time.Now()

	result := fakeClock{defaultNow}
	for _, opt := range opts {
		opt(&result)
	}
	return result
}

func (c fakeClock) Now() time.Time {
	return c.now
}

type FakeClockOption func(*fakeClock)

func WithNow(now time.Time) FakeClockOption {
	return func(clock *fakeClock) {
		clock.now = now
	}
}
