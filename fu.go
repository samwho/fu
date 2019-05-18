package fu

import (
	"context"
	"errors"
	"reflect"

	"github.com/samwho/fu/bifunction"
	"github.com/samwho/fu/filter"
	"github.com/samwho/fu/function"
	"github.com/samwho/fu/mapper"
	"github.com/samwho/fu/predicate"
	"github.com/samwho/fu/reducer"
)

func Map(ctx context.Context, f function.F, is []interface{}) ([]interface{}, error) {
	return mapper.New(f).Map(ctx, is)
}

func MapFn(ctx context.Context, f function.Fn, is []interface{}) ([]interface{}, error) {
	return mapper.NewFn(f).Map(ctx, is)
}

func Reduce(ctx context.Context, bf bifunction.B, i interface{}, is []interface{}) (interface{}, error) {
	return reducer.New(bf).Reduce(ctx, i, is)
}

func ReduceFn(ctx context.Context, bf bifunction.Fn, i interface{}, is []interface{}) (interface{}, error) {
	return reducer.NewFn(bf).Reduce(ctx, i, is)
}

func Filter(ctx context.Context, p predicate.P, is []interface{}) ([]interface{}, error) {
	return filter.New(p).Filter(ctx, is)
}

func FilterFn(ctx context.Context, p predicate.Fn, is []interface{}) ([]interface{}, error) {
	return filter.NewFn(p).Filter(ctx, is)
}

func Any(ctx context.Context, p predicate.P, is []interface{}) (bool, error) {
	for _, i := range is {
		b, err := p.Test(ctx, i)
		if err != nil {
			return false, err
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}

func AnyFn(ctx context.Context, p predicate.Fn, is []interface{}) (bool, error) {
	return Any(ctx, predicate.New(p), is)
}

func All(ctx context.Context, p predicate.P, is []interface{}) (bool, error) {
	for _, i := range is {
		b, err := p.Test(ctx, i)
		if err != nil {
			return false, err
		}
		if !b {
			return false, nil
		}
	}
	return true, nil
}

func AllFn(ctx context.Context, p predicate.Fn, is []interface{}) (bool, error) {
	return All(ctx, predicate.New(p), is)
}

func MapK(kf function.F) bifunction.B {
	return bifunction.New(func(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
		m, ok := i.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("type incompatible")
		}

		k, err := kf.Call(ctx, j)
		if err != nil {
			return nil, err
		}
		m[k] = j
		return m, nil
	})
}

func MapKV(kf function.F, vf function.F) bifunction.B {
	return bifunction.New(func(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
		m, ok := i.(map[interface{}]interface{})
		if !ok {
			return nil, errors.New("type incompatible")
		}

		k, err := kf.Call(ctx, j)
		if err != nil {
			return nil, err
		}
		v, err := vf.Call(ctx, j)
		if err != nil {
			return nil, err
		}
		m[k] = v
		return m, nil
	})
}

func GroupBy(ctx context.Context, f function.F, is []interface{}) (map[interface{}]interface{}, error) {
	r := reducer.New(MapK(f))
	m, err := r.Reduce(ctx, make(map[interface{}]interface{}), is)
	if err != nil {
		return nil, err
	}
	return m.(map[interface{}]interface{}), nil
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

func Sub() bifunction.B {
	return bifunction.New(
		func(ctx context.Context, a interface{}, b interface{}) (interface{}, error) {
			if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
				return 0, errors.New("types not compatible")
			}

			switch a.(type) {
			case int:
				return b.(int) - a.(int), nil
			case int32:
				return b.(int32) - a.(int32), nil
			case int64:
				return b.(int64) - a.(int64), nil
			case uint:
				return b.(uint) - a.(uint), nil
			case uint32:
				return b.(uint32) - a.(uint32), nil
			case uint64:
				return b.(uint64) - a.(uint64), nil
			case float32:
				return b.(float32) - a.(float32), nil
			case float64:
				return b.(float64) - a.(float64), nil
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
	return predicate.Or(Gt(a), Eq(a))
}

func Lte(a interface{}) predicate.P {
	return predicate.Or(Lt(a), Eq(a))
}

func Neq(a interface{}) predicate.P {
	return predicate.Not(Eq(a))
}
