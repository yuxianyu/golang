package main

import (
	"context"
	"fmt"
)

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}*/

func main() {
	fmt.Println(context.Background() == context.Background())
	fmt.Println(struct{}{} == struct{}{})
}
