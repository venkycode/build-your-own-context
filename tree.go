package byo_context

type children map[Context]struct{}

func (ch *children) addChild(ctx Context) {
	if *ch == nil {
		*ch = make(map[Context]struct{})
	}
	(*ch)[ctx] = struct{}{}
}

func (ch *children) removeChild(ctx Context) {
	if *ch == nil {
		panic("children: removeChild called on nil children")
	}
	delete(*ch, ctx)
}

func (ch *children) removeAll() {
	*ch = nil
}

type treeOps interface {
	addChild(ctx Context)
	removeChild(ctx Context)
	removeAll()
}

// assert that background implements the treeOps interface
var _ treeOps = &background{}

// assert that cancelCtx implements the treeOps interface
var _ treeOps = &cancelCtx{}

// assert that valueCtx implements the treeOps interface
var _ treeOps = &valueCtx{}

// assert that deadlineCtx implements the treeOps interface
var _ treeOps = &deadlineCtx{}
