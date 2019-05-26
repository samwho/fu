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

func TestInts(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5}
	actual, err := Ints(ctx, expected).Ints()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestInt32s(t *testing.T) {
	expected := []int32{0, 1, 2, 3, 4, 5}
	actual, err := Int32s(ctx, expected).Int32s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestInt64s(t *testing.T) {
	expected := []int64{0, 1, 2, 3, 4, 5}
	actual, err := Int64s(ctx, expected).Int64s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUints(t *testing.T) {
	expected := []uint{0, 1, 2, 3, 4, 5}
	actual, err := Uints(ctx, expected).Uints()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUint32s(t *testing.T) {
	expected := []uint32{0, 1, 2, 3, 4, 5}
	actual, err := Uint32s(ctx, expected).Uint32s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUint64s(t *testing.T) {
	expected := []uint64{0, 1, 2, 3, 4, 5}
	actual, err := Uint64s(ctx, expected).Uint64s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFloat32s(t *testing.T) {
	expected := []float32{0, 1, 2, 3, 4, 5}
	actual, err := Float32s(ctx, expected).Float32s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFloat64s(t *testing.T) {
	expected := []float64{0, 1, 2, 3, 4, 5}
	actual, err := Float64s(ctx, expected).Float64s()
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
