package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	mySleep(ctx, time.Duration(7*time.Second))
	cancel()
}

func mySleep(ctx context.Context, v time.Duration) {

	select {
	case <-time.After(v):
		fmt.Printf("this run after %d interval\n", v)
	case <-ctx.Done():
		fmt.Println("ctx done")
	}
}
