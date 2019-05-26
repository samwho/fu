package fu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapReduce(t *testing.T) {
	result, err :=
		Ints(ctx, []int{1, 2, 3, 4, 5}).
			Map(Add(1)).
			Map(String()).
			Reduce(Join(", "))

	require.NoError(t, err)
	assert.Equal(t, "2, 3, 4, 5, 6", result)
}
