package byo_context

import (
	"errors"
	"time"
)

var DeadlineExceeded = errors.New("context deadline exceeded")

type deadlineCtx struct {
	Context
	children
}

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	d := &deadlineCtx{
		Context: parent,
	}
	parent.(treeOps).addChild(d)

	return d, func() {
	}
}

func (d *deadlineCtx) cancel() {

}
