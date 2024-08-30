package byo_context

type valueCtx struct {
	Context
}

func WithValue(ctx Context, key, val interface{}) Context {
	return &valueCtx{
		Context: ctx,
	}
}
