package main

import "context"

func RunWithErrorHandling(ctx context.Context, fn func() error) error {
	err := fn()
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	return nil
}
