package funcutil

import (
	"context"
)

func Map(ctx context.Context, f Function, is []interface{}) ([]interface{}, error) {
	ret := make([]interface{}, len(is))
	for idx, i := range is {
		var err error
		ret[idx], err = f.Call(ctx, i)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func Reduce(ctx context.Context, bf BiFunction, is []interface{}) (interface{}, error) {
	if len(is) == 0 {
		return nil, nil
	}

	r := is[0]
	var err error
	for i := 1; i < len(is); i++ {
		r, err = bf.Call(ctx, r, is[i])
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func Filter(ctx context.Context, p Predicate, is []interface{}) ([]interface{}, error) {
	var filtered []interface{}
	for _, i := range is {
		b, err := p.Test(ctx, i)
		if err != nil {
			return nil, err
		}
		if !b {
			continue
		}
		filtered = append(filtered, i)
	}
	return filtered, nil
}
