package main

import (
	"fmt"
	"time"
)

func doneUsecase() {
	go hearbeat()

	go pushStateUpdates()

	mainWork()
}

func hearbeat() {
	for {
		onebeat()
	}
}

func pushStateUpdates() {
	for {
		onePushStateUpdate()
	}
}

func mainWork() {
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
