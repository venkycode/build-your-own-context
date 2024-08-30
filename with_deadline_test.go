package byo_context

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func after(dur time.Duration) time.Time {
	return time.Now().Add(dur)
}

func Test_WithDeadline(t *testing.T) {
	root := Background()

	deadline := after(100 * time.Millisecond)

	ctx, cancel := WithDeadline(root, deadline)

	time.Sleep(50 * time.Millisecond)

	select {
	case <-ctx.Done():
		t.Error("context should not be done")
	default:
	}
	exprectDeadline(t, ctx, deadline)
	time.Sleep(50 * time.Millisecond)

	select {
	case <-ctx.Done():
	default:
		t.Error("context should be done")
	}
	exprectDeadline(t, ctx, deadline)

	// no effect of cancel after deadline
	cancel()

	select {
	case <-ctx.Done():
	default:
		t.Error("context should be done")
	}
	exprectDeadline(t, ctx, deadline)

	require.ErrorIs(t, ctx.Err(), DeadlineExceeded)

}

func exprectDeadline(t *testing.T, ctx Context, deadline time.Time) {
	t.Helper()
	d, ok := ctx.Deadline()
	require.True(t, ok)
	require.Equal(t, deadline, d)
}

func Test_Deadline_override_deadline(t *testing.T) {
	root := Background()

	deadline := after(100 * time.Millisecond)

	ctx1, _ := WithDeadline(root, deadline)

	newDeadline := after(200 * time.Millisecond)

	ctx2, _ := WithDeadline(ctx1, newDeadline)

	exprectDeadline(t, ctx2, newDeadline)
	exprectDeadline(t, ctx1, deadline)
}

func Test_Deadline_early_cancel(t *testing.T) {
	root := Background()

	deadline := after(100 * time.Millisecond)

	ctx, cancel := WithDeadline(root, deadline)

	cancel()

	select {
	case <-ctx.Done():
	default:
		t.Error("context should be done")
	}

	require.ErrorIs(t, ctx.Err(), Canceled)
}
