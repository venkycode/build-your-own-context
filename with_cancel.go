package byo_context

import (
	"errors"
)

// not concurrent safe
type cancelCtx struct {
	Context // embed the parent context (Deadline, Value) derived from the parent context
	children
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

// when this function is called, the context is cancelled
type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}
	parent.(treeOps).addChild(c)
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
	cancelChildren(c.children)
}

func cancelChildren(children children) {
	for child := range children {
		cancellable, ok := child.(interface{ cancel() })
		if ok {
			cancellable.cancel()
		}
	}
	children.removeAll()
}
