package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := time.Now()
	sig1 := sig(10 * time.Second)
	sig2 := sig(2 * time.Second)
	sig3 := sig(3 * time.Second)
	orCh := or(sig1, sig2, sig3)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		case <-orCh:
			fmt.Println("or channel closed because of one input chan was closed")
			cancel()
		}
	}(ctx)

	<-ctx.Done()

	fmt.Printf("done after:%v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		cases := make([]reflect.SelectCase, len(channels))
		for i, channel := range channels {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channel),
			}
		}

		chosen, _, ok := reflect.Select(cases)
		if !ok {
			fmt.Printf("channel: %d is done ", chosen)
		}
	}()

	return ch
}

func sig(duration time.Duration) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		time.Sleep(duration)
	}()

	return ch
}
