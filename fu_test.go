package fu

import (
	"context"
	"testing"

	"github.com/samwho/fu/function"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func TestReduce(t *testing.T) {
	nums := []interface{}{1, 2, 3, 4, 5}
	sum, err := Reduce(ctx, Sum(), 0, nums)
	require.NoError(t, err)
	assert.Equal(t, 15, sum)
}

func TestGroupBy(t *testing.T) {
	type record struct {
		id   int
		data string
	}

	rs := []interface{}{
		record{id: 1, data: "hello"},
		record{id: 2, data: "world"},
	}

	f := function.New(func(ctx context.Context, i interface{}) (interface{}, error) {
		return i.(record).id, nil
	})

	m, err := GroupBy(ctx, f, rs)
	require.NoError(t, err)

	assert.Equal(t, m[1], rs[0])
	assert.Equal(t, m[2], rs[1])
}
