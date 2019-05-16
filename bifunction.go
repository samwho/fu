package funcutil

import (
	"context"
	"errors"
	"reflect"
)

type BiFunction interface {
	Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error)
}

type BiFunctionFn func(ctx context.Context, i interface{}, j interface{}) (interface{}, error)

type bifunctionImpl struct {
	bf BiFunctionFn
}

func (bf *bifunctionImpl) Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
	return bf.bf(ctx, i, j)
}

func NewBiFunction(bf BiFunctionFn) BiFunction {
	return &bifunctionImpl{bf: bf}
}

type multiBiFn struct {
	bfs []BiFunction
}

func (mbf *multiBiFn) Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
	for _, bf := range mbf.bfs {
		var err error
		i, err = bf.Call(ctx, i, j)
		if err != nil {
			return nil, err
		}
	}
	return i, nil
}

func Apply(i interface{}, bf BiFunction) Function {
	return NewFunction(func(ctx context.Context, j interface{}) (interface{}, error) {
		return bf.Call(ctx, i, j)
	})
}

func Sum() BiFunction {
	return NewBiFunction(
		func(ctx context.Context, a interface{}, b interface{}) (interface{}, error) {
			if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
				return 0, errors.New("types not compatible")
			}

			switch a.(type) {
			case int:
				return b.(int) + a.(int), nil
			case int32:
				return b.(int32) + a.(int32), nil
			case int64:
				return b.(int64) + a.(int64), nil
			case uint:
				return b.(uint) + a.(uint), nil
			case uint32:
				return b.(uint32) + a.(uint32), nil
			case uint64:
				return b.(uint64) + a.(uint64), nil
			case float32:
				return b.(float32) + a.(float32), nil
			case float64:
				return b.(float64) + a.(float64), nil
			default:
				return false, errors.New("types not comparable")
			}
		})
}
