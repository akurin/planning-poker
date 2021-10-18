package systemclock

import (
	"backend/internal/domain"
	"time"
)

type clock struct{}

func New() domain.Clock {
	return clock{}
}

func (c clock) Now() time.Time {
	return time.Now()
}
