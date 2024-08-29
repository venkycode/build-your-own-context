package byo_context

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	require.Nil(t, ctx.Err())

	_, ok := ctx.Deadline()
	require.False(t, ok)

	require.NotNil(t, ctx.Done())

	require.Nil(t, ctx.Value(nil))

	cancel()
	require.ErrorIs(t, ctx.Err(), Canceled)

	// test that cancel() is idempotent
	cancel()
	require.ErrorIs(t, ctx.Err(), Canceled)

	select {
	case <-ctx.Done():
		t.Log("ctx.Done() is closed after cancel()")
	default:
		t.Error("ctx.Done() should be closed after cancel()")
	}
}

func Test_WithCancel_propogation(t *testing.T) {
	parent, cancel := WithCancel(Background())
	child, _ := WithCancel(parent)

	require.Nil(t, child.Err())

	cancel()
	require.ErrorIs(t, child.Err(), Canceled)
}

func Test_WithCancel_propogation2(t *testing.T) {
	/*
		Creates a context tree like this:
		Background-root
		|
		|---cancelCtx1---cancelCtx3---cancelCtx5
		|		|
		|		|---cancelCtx4
		|
		|
		|---cancelCtx2---cancelCtx6

	*/

	root := Background()

	cancelCtx1, cancel1 := WithCancel(root)

	cancelCtx2, _ := WithCancel(root)

	cancelCtx3, _ := WithCancel(cancelCtx1)

	cancelCtx4, _ := WithCancel(cancelCtx1)

	cancelCtx5, _ := WithCancel(cancelCtx3)

	cancelCtx6, cancel6 := WithCancel(cancelCtx2)

	cancel1()

	require.ErrorIs(t, cancelCtx1.Err(), Canceled)
	require.ErrorIs(t, cancelCtx3.Err(), Canceled)
	require.ErrorIs(t, cancelCtx5.Err(), Canceled)
	require.ErrorIs(t, cancelCtx4.Err(), Canceled)

	require.Nil(t, cancelCtx2.Err())

	cancel6()
	require.ErrorIs(t, cancelCtx6.Err(), Canceled)
	require.Nil(t, cancelCtx2.Err())
}
