package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	defer cancel()
	go performTask(ctx)

	select {
	// case <- ctx.Value()
	case <-ctx.Done():
		fmt.Println("TImed out")
		return
	}

}

func performTask(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("Time out")
	case <-time.Tick(time.Microsecond * 300):
		fmt.Println("doing something good.")
	}

}
