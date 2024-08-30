package byo_context

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WithValue(t *testing.T) {
	/*
		ctx tree structure:

			root
			|
			ctx1([k1, v1]) - ctx3- [k1,v1, k3, v3] - ctx4 ([k1,vv1], [k3, v3])
			|
			ctx2([k1, v1], [k2, v2]) - ctx5 ([k1,v1],[k2, v2], [k5, v5])
			|
			ctx6([k1, v1], [k2, vv2])

	*/

	k1 := "key1"
	v1 := "value1"

	root := Background()

	ctx1 := WithValue(root, k1, v1)

	k2 := "key2"
	v2 := "value2"

	ctx2 := WithValue(ctx1, k2, v2)

	k3 := "key3"
	v3 := "value3"

	ctx3 := WithValue(ctx1, k3, v3)

	vv1 := "value1-2"

	ctx4 := WithValue(ctx3, k1, vv1)

	k5 := "key5"
	v5 := "value5"

	ctx5 := WithValue(ctx2, k5, v5)

	vv2 := "value2-2"
	ctx6 := WithValue(ctx2, k2, vv2)

	require.Equal(t, v1, ctx1.Value(k1))
	require.Nil(t, ctx1.Value(k2))

	require.Equal(t, v1, ctx2.Value(k1))
	require.Equal(t, v2, ctx2.Value(k2))

	require.Equal(t, v1, ctx3.Value(k1))
	require.Nil(t, ctx3.Value(k2))
	require.Equal(t, v3, ctx3.Value(k3))

	require.Equal(t, vv1, ctx4.Value(k1))
	require.Nil(t, ctx4.Value(k2))
	require.Equal(t, v3, ctx4.Value(k3))
	require.Equal(t, vv1, ctx4.Value(k1))

	require.Equal(t, v1, ctx5.Value(k1))
	require.Equal(t, v2, ctx5.Value(k2))
	require.Nil(t, ctx5.Value(k3))
	require.Equal(t, v5, ctx5.Value(k5))

	require.Equal(t, v1, ctx6.Value(k1))
	require.Equal(t, vv2, ctx6.Value(k2))
	require.Nil(t, ctx6.Value(k3))
}
