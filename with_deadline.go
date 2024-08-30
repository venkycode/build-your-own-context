package byo_context

import (
	"errors"
	"time"
)

var DeadlineExceeded = errors.New("context deadline exceeded")

type deadlineCtx struct {
	cancelCtx // embed the cancelCtx to get the cancel() method
	deadline  time.Time
}

// func (c *deadlineCtx) Done() <-chan struct{}  derived from cancelCtx
// func (c *deadlineCtx) Err() error  derived from cancelCtx

func (c *deadlineCtx) Deadline() (time.Time, bool) {
	return c.deadline, true
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {

	if cur, ok := parent.Deadline(); ok && cur.Before(deadline) {
		// parent has a deadline that is sooner than the deadline
		// keep parent deadline
		return WithCancel(parent)
	}
	c := &deadlineCtx{
		cancelCtx: embedCancelCtx(parent),
	}
	parent.(treeOps).addChild(c)

	c.deadline = deadline
	dur := time.Until(deadline)

	if dur <= 0 {
		// the deadline has already passed
		c.cancel(DeadlineExceeded)

		return c, func() {}
	}

	time.AfterFunc(dur, func() {
		c.cancel(DeadlineExceeded) // cancel the context after the deadline has passed
	})

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
