package cloud_patterns

import (
	"context"
	"errors"
	"sync"
	"time"
)

var TooManyCallsErr = errors.New("too many calls")

type throttle struct {
	tokens   uint
	max      uint
	refill   uint
	duration time.Duration
	once     sync.Once
}

func NewThrottle() *throttle {
	return &throttle{}
}

// Execute is missing goroutine safety
func (t *throttle) Execute(ctx context.Context, effector Effector) (Value, error) {
	if ctx.Err() != nil {
		return Value(""), ctx.Err()
	}

	t.once.Do(func() {
		ticket := time.NewTicker(t.duration)

		go func() {
			defer ticket.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticket.C:
					tokens := t.tokens + t.refill
					if tokens > t.max {
						tokens = t.max
					}
					t.tokens = tokens
				}
			}
		}()
	})

	if t.tokens <= 0 {
		return Value(""), TooManyCallsErr
	}

	t.tokens--

	return effector(ctx)
}
