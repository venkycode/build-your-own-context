package byo_context

type valueCtx struct {
	Context
	children
	key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if key == c.key {
		return c.val
	}
	return c.Context.Value(key)
}

func WithValue(parent Context, key, val interface{}) Context {

	valueCtx := &valueCtx{
		Context: parent,
		key:     key,
		val:     val,
	}

	parent.(treeOps).addChild(valueCtx)
	return parent
}
