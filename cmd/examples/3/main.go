package main

import (
	context "github.com/venkycode/build-your-own-context"
)

func main() {
	ctx := context.Background()

	deadlineUseCase(ctx)
	busyWork(ctx)
}

func busyWork(ctx context.Context) {
	for {
	}
}
