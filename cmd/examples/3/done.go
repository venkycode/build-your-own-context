package main

import (
	"fmt"
	"time"
)

func doneUsecase() {
	done := make(chan struct{})
	go hearbeat(done)

	go pushStateUpdates(done)

	mainWork(done)
}

func hearbeat(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("------>stopped heartbeat<----")
			return
		default:
			onebeat()
		}
	}
}

func pushStateUpdates(done chan struct{}) {

	for {
		select {
		case <-done:
			fmt.Println("------>stopped state updates<----")
			return
		default:
			onePushStateUpdate()
		}
	}
}

func mainWork(done chan struct{}) {
	workFor(2000)
	fmt.Println("->Main work done<-")
	close(done)
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
