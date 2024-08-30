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

	deadline1 := after(100 * time.Millisecond)

	ctx1, _ := WithDeadline(root, deadline1)

	// new deadline is before the parent deadline
	deadline2 := after(50 * time.Millisecond)

	ctx2, _ := WithDeadline(ctx1, deadline2)

	exprectDeadline(t, ctx2, deadline2)
	exprectDeadline(t, ctx1, deadline1)

	dealine3 := after(150 * time.Millisecond)
	ctx3, _ := WithDeadline(ctx1, dealine3) // dedline is after the parent deadline
	exprectDeadline(t, ctx3, deadline1)     // deadline is the parent deadline, as it is before the new deadline
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
