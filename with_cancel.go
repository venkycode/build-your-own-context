package byo_context

import (
	"errors"
	"time"
)

// not concurrent safe
type cancelCtx struct {
	done chan struct{}
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}

var Canceled = errors.New("context canceled")

func (c *cancelCtx) Err() error {
	select {
	case <-c.done:
		return Canceled
	default:
		return nil
	}
}

func (c *cancelCtx) Deadline() (deadline time.Time, ok bool) {
	//todo implement
	return
}

func (c *cancelCtx) Value(key interface{}) interface{} {
	//todo implement
	return nil
}

// when this function is called, the context is cancelled
type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := &cancelCtx{
		done: make(chan struct{}),
	}
	return c, func() {
		c.cancel()
	}
}

func (c *cancelCtx) cancel() {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
}
