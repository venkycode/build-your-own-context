package byo_context

import (
	"errors"
)

// not concurrent safe
type cancelCtx struct {
	Context // embed the parent context (Deadline, Value) derived from the parent context
	children
	done chan struct{}
	err  error
}

func (c *cancelCtx) Done() <-chan struct{} {
	return c.done
}

var Canceled = errors.New("context canceled")

func (c *cancelCtx) Err() error {
	return c.err
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
		c.cancel(Canceled)
	}
}

func (c *cancelCtx) cancel(err error) {
	select {
	case <-c.done:
	default:
		close(c.done)
		cancelChildren(c.children, err)
		c.err = err
	}
}

func cancelChildren(children children, err error) {
	for child := range children {
		canceller, ok := child.(cannceller)
		if ok {
			canceller.cancel(err)
		}
	}
	children.removeAll()
}

type cannceller interface {
	cancel(error)
}
