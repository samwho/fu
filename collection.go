package fu

import (
	"context"
	"errors"

	"github.com/samwho/fu/predicate"

	"github.com/samwho/fu/bifunction"
	"github.com/samwho/fu/function"
)

type Collection struct {
	ctx context.Context
	is  []interface{}
	err error
}

func (c *Collection) Error() error {
	return c.err
}

func (c *Collection) MapFn(f function.Fn) *Collection {
	return c.Map(function.New(f))
}

func (c *Collection) Map(f function.F) *Collection {
	if c.err != nil {
		return c
	}
	c.is, c.err = Map(c.ctx, c.is, f)
	return c
}

func (c *Collection) ParallelMapFn(parallelism int, f function.Fn) *Collection {
	return c.ParallelMap(parallelism, function.New(f))
}

func (c *Collection) ParallelMap(parallelism int, f function.F) *Collection {
	if c.err != nil {
		return c
	}
	c.is, c.err = ParallelMap(c.ctx, parallelism, c.is, f)
	return c
}

func (c *Collection) SelectFn(p predicate.Fn) *Collection {
	return c.Select(predicate.New(p))
}

func (c *Collection) Select(p predicate.P) *Collection {
	if c.err != nil {
		return c
	}
	c.is, c.err = Select(c.ctx, c.is, p)
	return c
}

func (c *Collection) RejectFn(p predicate.Fn) *Collection {
	return c.Reject(predicate.New(p))
}

func (c *Collection) Reject(p predicate.P) *Collection {
	if c.err != nil {
		return c
	}
	c.is, c.err = Reject(c.ctx, c.is, p)
	return c
}

func (c *Collection) ReduceFn(bf bifunction.Fn) (interface{}, error) {
	return c.Reduce(bifunction.New(bf))
}

func (c *Collection) Reduce(bf bifunction.B) (interface{}, error) {
	if c.err != nil {
		return nil, c.err
	}
	var i interface{}
	i, c.err = Reduce(c.ctx, c.is, bf)
	return i, c.err
}

func Ints(ctx context.Context, in []int) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Int32s(ctx context.Context, in []int32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Int64s(ctx context.Context, in []int64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Uints(ctx context.Context, in []uint) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Uint32s(ctx context.Context, in []uint32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Uint64s(ctx context.Context, in []uint64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Float32s(ctx context.Context, in []float32) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Float64s(ctx context.Context, in []float64) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func Strings(ctx context.Context, in []string) *Collection {
	is := make([]interface{}, 0, len(in))
	for _, i := range in {
		is = append(is, i)
	}
	return &Collection{ctx, is, nil}
}

func (c *Collection) Ints() ([]int, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]int, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(int)
		if !ok {
			return nil, errors.New("not all elements are type int")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Int32s() ([]int32, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]int32, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(int32)
		if !ok {
			return nil, errors.New("not all elements are type int32")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Int64s() ([]int64, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]int64, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(int64)
		if !ok {
			return nil, errors.New("not all elements are type int64")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Uints() ([]uint, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]uint, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(uint)
		if !ok {
			return nil, errors.New("not all elements are type uint")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Uint32s() ([]uint32, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]uint32, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(uint32)
		if !ok {
			return nil, errors.New("not all elements are type uint32")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Uint64s() ([]uint64, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]uint64, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(uint64)
		if !ok {
			return nil, errors.New("not all elements are type uint64")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Float32s() ([]float32, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]float32, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(float32)
		if !ok {
			return nil, errors.New("not all elements are type float32")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Float64s() ([]float64, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]float64, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(float64)
		if !ok {
			return nil, errors.New("not all elements are type float64")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Strings() ([]string, error) {
	if c.err != nil {
		return nil, c.err
	}
	ret := make([]string, 0, len(c.is))
	for _, i := range c.is {
		e, ok := i.(string)
		if !ok {
			return nil, errors.New("not all elements are type string")
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func (c *Collection) Interfaces() ([]interface{}, error) {
	return c.is, c.err
}
