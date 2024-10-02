package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	g, gctx := errgroup.WithContext(ctx)

	var a string
	g.Go(func() error {
		return RunWithErrorHandling(gctx, func() error {
			var err error
			a, err = sampleA(gctx)
			return err
		})
	})

	var b int
	g.Go(func() error {
		return RunWithErrorHandling(gctx, func() error {
			var err error
			b, err = sampleB(gctx)
			return err
		})
	})

	err := g.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Println(a, b)
}


func sampleA(ctx context.Context) (string, error) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("abc")
			return "abc", nil
		case <-ctx.Done():
			if ctx.Err() != nil {
				return "", ctx.Err()
			}
		}
	}
}

func sampleB(ctx context.Context) (int, error) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("b")

	return 0, nil
}
