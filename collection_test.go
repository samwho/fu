package fu

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectionReduce(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).Reduce(Join(", "))
	require.NoError(t, err)
	assert.Equal(t, "hello, world", result)
}

func TestCollectionReduceAfterErr(t *testing.T) {
	_, err := Strings(ctx, []string{"hello", "world"}).Map(Add(1)).Reduce(Join(", "))
	assert.Error(t, err)
}

func TestCollectionReduceFn(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).ReduceFn(func(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
		return strings.Join([]string{i.(string), j.(string)}, ", "), nil
	})
	require.NoError(t, err)
	assert.Equal(t, "hello, world", result)
}

func TestCollectionMap(t *testing.T) {
	result, err := Ints(ctx, []int{1, 2, 3, 4, 5}).Map(Add(1)).Ints()
	require.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, result)
}

func TestCollectionMapAfterErr(t *testing.T) {
	_, err := Strings(ctx, []string{"hello, world"}).Map(Add(1)).Map(Add(1)).Ints()
	assert.Error(t, err)
}

func TestCollectionParallelMap(t *testing.T) {
	result, err := Ints(ctx, []int{1, 2, 3, 4, 5}).ParallelMap(8, Add(1)).Ints()
	require.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, result)
}

func TestCollectionParallelMapAfterErr(t *testing.T) {
	_, err := Strings(ctx, []string{"hello, world"}).Map(Add(1)).ParallelMap(8, Add(1)).Ints()
	assert.Error(t, err)
}

func TestCollectionMapFn(t *testing.T) {
	result, err := Ints(ctx, []int{1, 2, 3, 4, 5}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	}).Ints()
	require.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, result)
}

func TestCollectionParallelMapFn(t *testing.T) {
	result, err := Ints(ctx, []int{1, 2, 3, 4, 5}).ParallelMapFn(8, func(ctx context.Context, i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	}).Ints()
	require.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, result)
}

func TestCollectionError(t *testing.T) {
	c := Strings(ctx, []string{"hello, world"}).Map(Add(1))
	assert.Error(t, c.Error())
}

func TestCollectionSelect(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).Select(Gt("jam")).Strings()
	require.NoError(t, err)
	assert.Equal(t, []string{"world"}, result)
}

func TestCollectionSelectFn(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).SelectFn(func(ctx context.Context, i interface{}) (bool, error) {
		return i.(string) > "jam", nil
	}).Strings()
	require.NoError(t, err)
	assert.Equal(t, []string{"world"}, result)
}

func TestCollectionReject(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).Reject(Gt("jam")).Strings()
	require.NoError(t, err)
	assert.Equal(t, []string{"hello"}, result)
}

func TestCollectionRejectFn(t *testing.T) {
	result, err := Strings(ctx, []string{"hello", "world"}).RejectFn(func(ctx context.Context, i interface{}) (bool, error) {
		return i.(string) > "jam", nil
	}).Strings()
	require.NoError(t, err)
	assert.Equal(t, []string{"hello"}, result)
}

func TestCollectionSelectErr(t *testing.T) {
	_, err := Strings(ctx, []string{"hello", "world"}).Map(Add(1)).Select(Gt("jam")).Strings()
	assert.Error(t, err)
}

func TestCollectionRejectErr(t *testing.T) {
	_, err := Strings(ctx, []string{"hello", "world"}).Map(Add(1)).Reject(Gt("jam")).Strings()
	assert.Error(t, err)
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

func TestInterfaces(t *testing.T) {
	actual, err := Float64s(ctx, []float64{0}).Interfaces()
	require.NoError(t, err)
	assert.Equal(t, []interface{}{float64(0)}, actual)
}

func TestIntsError(t *testing.T) {
	_, err := Ints(ctx, []int{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Ints()
	assert.Error(t, err)
}

func TestInt32sError(t *testing.T) {
	_, err := Int32s(ctx, []int32{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Int32s()
	assert.Error(t, err)
}

func TestInt64sError(t *testing.T) {
	_, err := Int64s(ctx, []int64{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Int64s()
	assert.Error(t, err)
}

func TestUintsError(t *testing.T) {
	_, err := Uints(ctx, []uint{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Uints()
	assert.Error(t, err)
}

func TestUint32sError(t *testing.T) {
	_, err := Uint32s(ctx, []uint32{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Uint32s()
	assert.Error(t, err)
}

func TestUint64sError(t *testing.T) {
	_, err := Uint64s(ctx, []uint64{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Uint64s()
	assert.Error(t, err)
}

func TestFloat32sError(t *testing.T) {
	_, err := Float32s(ctx, []float32{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Float32s()
	assert.Error(t, err)
}

func TestFloat64sError(t *testing.T) {
	_, err := Float64s(ctx, []float64{0, 1}).MapFn(func(ctx context.Context, i interface{}) (interface{}, error) {
		return nil, errors.New("")
	}).Float64s()
	assert.Error(t, err)
}

func TestInvalidTypeConversions(t *testing.T) {
	var err error
	_, err = Strings(ctx, []string{"hello"}).Ints()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Int32s()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Int64s()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Uints()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Uint32s()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Uint64s()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Float32s()
	assert.Error(t, err)
	_, err = Strings(ctx, []string{"hello"}).Float64s()
	assert.Error(t, err)

	_, err = Ints(ctx, []int{0}).Strings()
	assert.Error(t, err)
}
