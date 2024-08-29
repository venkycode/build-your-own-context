package byo_context

import "time"

// this is usually the root of a context tree
// it is never (or can not be) cancelled, has no values, and has no deadline
// usually initialized in main() and passed down topmost function calls
type background struct {
	children
}

func (b *background) Value(key interface{}) interface{} {
	return nil
}

func (b *background) Deadline() (deadline time.Time, ok bool) {
	return
}

func (b *background) Done() <-chan struct{} {
	return nil
}

func (b *background) Err() error {
	return nil
}

func Background() Context {
	return &background{}
}
