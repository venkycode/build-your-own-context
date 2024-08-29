package byo_context

import "fmt"

type children map[Context]struct{}

func (ch *children) addChild(ctx Context) {
	if *ch == nil {
		*ch = make(map[Context]struct{})
	}
	fmt.Println("children: addChild called")
	(*ch)[ctx] = struct{}{}
}

func (ch *children) removeChild(ctx Context) {
	if *ch == nil {
		panic("children: removeChild called on nil children")
	}
	delete(*ch, ctx)
}

type treeOps interface {
	addChild(ctx Context)
	removeChild(ctx Context)
}

// assert that background implements the treeOps interface
var _ treeOps = &background{}

// assert that cancelCtx implements the treeOps interface
var _ treeOps = &cancelCtx{}
