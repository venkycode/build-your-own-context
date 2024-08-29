package byo_context

import (
	"context"
	"time"
)

// asserts that byo_context.Context is compatible with context.Context
var _ Context = context.Context(nil)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
