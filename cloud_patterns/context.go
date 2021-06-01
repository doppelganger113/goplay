package cloud_patterns

import (
	"context"
	"time"
)

type Value []byte

func SlowOperation(ctx context.Context, seconds uint) (Value, error) {
	select {
	case <-time.After(time.Second * time.Duration(seconds)):
		return []byte("Completed hard work."), nil
	case <-ctx.Done():
		return []byte(""), ctx.Err()
	}
}

func Stream(ctx context.Context, out chan<- Value) error {
	derivedCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := SlowOperation(derivedCtx, 5)
	if err != nil {
		return err
	}

	for {
		select {
		case out <- res:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
