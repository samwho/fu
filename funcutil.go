package funcutil

import (
	"context"
	"errors"
	"reflect"

	"github.com/samwho/funcutil/bifunction"
	"github.com/samwho/funcutil/function"
	"github.com/samwho/funcutil/predicate"
)

func Map(ctx context.Context, f function.F, is []interface{}) ([]interface{}, error) {
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

func Reduce(ctx context.Context, bf bifunction.B, is []interface{}) (interface{}, error) {
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

func Filter(ctx context.Context, p predicate.P, is []interface{}) ([]interface{}, error) {
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

func Sum() bifunction.B {
	return bifunction.New(
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

func Gt(a interface{}) predicate.P {
	return predicate.New(func(ctx context.Context, b interface{}) (bool, error) {
		if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
			return false, errors.New("types not compatible")
		}

		switch a.(type) {
		case int:
			return b.(int) > a.(int), nil
		case int32:
			return b.(int32) > a.(int32), nil
		case int64:
			return b.(int64) > a.(int64), nil
		case uint:
			return b.(uint) > a.(uint), nil
		case uint32:
			return b.(uint32) > a.(uint32), nil
		case uint64:
			return b.(uint64) > a.(uint64), nil
		case float32:
			return b.(float32) > a.(float32), nil
		case float64:
			return b.(float64) > a.(float64), nil
		case string:
			return b.(string) > a.(string), nil
		default:
			return false, errors.New("types not comparable")
		}
	})
}

func Lt(a interface{}) predicate.P {
	return predicate.New(func(ctx context.Context, b interface{}) (bool, error) {
		if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
			return false, errors.New("types not compatible")
		}

		switch a.(type) {
		case int:
			return b.(int) < a.(int), nil
		case int32:
			return b.(int32) < a.(int32), nil
		case int64:
			return b.(int64) < a.(int64), nil
		case uint:
			return b.(uint) < a.(uint), nil
		case uint32:
			return b.(uint32) < a.(uint32), nil
		case uint64:
			return b.(uint64) < a.(uint64), nil
		case float32:
			return b.(float32) < a.(float32), nil
		case float64:
			return b.(float64) < a.(float64), nil
		case string:
			return b.(string) < a.(string), nil
		default:
			return false, errors.New("types not comparable")
		}
	})
}

func Eq(a interface{}) predicate.P {
	return predicate.New(func(ctx context.Context, b interface{}) (bool, error) {
		return reflect.DeepEqual(a, b), nil
	})
}

func Gte(a interface{}) predicate.P {
	return predicate.Any(Gt(a), Eq(a))
}

func Lte(a interface{}) predicate.P {
	return predicate.Any(Lt(a), Eq(a))
}

func Neq(a interface{}) predicate.P {
	return predicate.Not(Eq(a))
}
