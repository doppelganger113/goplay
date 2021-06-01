package cloud_patterns

import (
	"context"
	"math/rand"
	"time"
)

type Effector func(ctx context.Context) (Value, error)

type retry struct {
	retries uint
	delay   time.Duration
}

func NewRetry(delay time.Duration) *retry {
	return &retry{delay: delay}
}

// Execute will retry failed request with a jitter backoff algorithm. Please
// execute a rand.Seed(time.Now().UTC().UnixNano()) at the start of the main
// function to have different numbers generated
func (r *retry) Execute(ctx context.Context, effector Effector) (Value, error) {
	response, err := effector(ctx)
	baseBackoff, maximumBackoff := time.Second, time.Minute

	for backoff := baseBackoff; err != nil; backoff <<= 1 {
		if backoff > maximumBackoff {
			backoff = maximumBackoff
		}

		jitter := rand.Int63n(int64(backoff * 3))
		sleep := baseBackoff + time.Duration(jitter)
		time.Sleep(sleep)
		response, err = effector(ctx)
	}

	return response, err
}
