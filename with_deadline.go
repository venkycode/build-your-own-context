package byo_context

import (
	"errors"
	"time"
)

var DeadlineExceeded = errors.New("context deadline exceeded")

type deadlineCtx struct {
	cancelCtx // embed the cancelCtx to get the cancel() method
}

// func (c *deadlineCtx) Done() <-chan struct{}  derived from cancelCtx
// func (c *deadlineCtx) Err() error  derived from cancelCtx

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	c := &deadlineCtx{
		cancelCtx: embedCancelCtx(parent),
	}
	parent.(treeOps).addChild(c)

	return c, func() {
		c.cancel(Canceled)
	}
}

func embedCancelCtx(parent Context) cancelCtx {
	return cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}
}
