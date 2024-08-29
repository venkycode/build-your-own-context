package main

import (
	"context"
	"fmt"
)

// Cancelling a context does not ensure that the fuctions receiving that context suddenly stop
func main() {
	ctx1 := context.Background()
	ctx2 := context.WithValue(ctx1, "k1", "v1")

	fmt.Println("value of 'k1' in ctx1:", ctx1.Value("k1"))
	fmt.Println("value of 'k1' in ctx2:", ctx2.Value("k1"))
}
