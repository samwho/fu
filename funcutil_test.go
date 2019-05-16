package funcutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func TestSum(t *testing.T) {
	nums := []interface{}{0, 1, 2, 3, 4, 5}
	sum, err := Reduce(ctx, Sum(), nums)
	require.NoError(t, err)
	assert.Equal(t, 15, sum)
}
