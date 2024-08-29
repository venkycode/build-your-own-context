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
