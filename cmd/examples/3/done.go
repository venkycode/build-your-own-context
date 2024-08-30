package main

import (
	"fmt"
	"time"

	context "github.com/venkycode/build-your-own-context"
)

func withCancelUsecase(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go hearbeat(ctx)

	go pushStateUpdates(ctx)

	mainWork(ctx)
}

func hearbeat(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("------>stopped heartbeat<----")
			return
		default:
			onebeat()
		}
	}
}

func pushStateUpdates(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("------>stopped state updates<----")
			return
		default:
			onePushStateUpdate()
		}
	}
}

func mainWork(ctx context.Context) {
	workFor(2000)
	fmt.Println("->Main work done<-")
}

func onebeat() {
	workFor(100)
	fmt.Println("one beat check done")
}

func onePushStateUpdate() {
	workFor(200)
	fmt.Println("one state update push done")
}

func workFor(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
