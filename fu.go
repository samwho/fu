package fu

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/samwho/fu/bifunction"
	"github.com/samwho/fu/filter"
	"github.com/samwho/fu/function"
	"github.com/samwho/fu/mapper"
	"github.com/samwho/fu/predicate"
	"github.com/samwho/fu/reducer"
)

func Map(ctx context.Context, is []interface{}, f function.F) ([]interface{}, error) {
	return mapper.New(f).Map(ctx, is)
}

func MapFn(ctx context.Context, is []interface{}, f function.Fn) ([]interface{}, error) {
	return mapper.NewFn(f).Map(ctx, is)
}

func ParallelMap(ctx context.Context, paralellism int, is []interface{}, f function.F) ([]interface{}, error) {
	return mapper.Parallel(paralellism, f).Map(ctx, is)
}

func ParallelMapFn(ctx context.Context, paralellism int, is []interface{}, f function.Fn) ([]interface{}, error) {
	return mapper.Parallel(paralellism, function.New(f)).Map(ctx, is)
}

func Reduce(ctx context.Context, i interface{}, is []interface{}, bf bifunction.B) (interface{}, error) {
	return reducer.New(bf).Reduce(ctx, i, is)
}

func ReduceFn(ctx context.Context, i interface{}, is []interface{}, bf bifunction.Fn) (interface{}, error) {
	return reducer.NewFn(bf).Reduce(ctx, i, is)
}

func Select(ctx context.Context, is []interface{}, p predicate.P) ([]interface{}, error) {
	return filter.New(p).Filter(ctx, is)
}

func SelectFn(ctx context.Context, is []interface{}, p predicate.Fn) ([]interface{}, error) {
	return filter.NewFn(p).Filter(ctx, is)
}

func Reject(ctx context.Context, is []interface{}, p predicate.P) ([]interface{}, error) {
	return filter.New(Not(p)).Filter(ctx, is)
}

func RejectFn(ctx context.Context, is []interface{}, p predicate.Fn) ([]interface{}, error) {
	return filter.New(Not(predicate.New(p))).Filter(ctx, is)
}

func Any(ctx context.Context, is []interface{}, p predicate.P) (bool, error) {
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

func AnyFn(ctx context.Context, is []interface{}, p predicate.Fn) (bool, error) {
	return Any(ctx, is, predicate.New(p))
}

func All(ctx context.Context, is []interface{}, p predicate.P) (bool, error) {
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

func AllFn(ctx context.Context, is []interface{}, p predicate.Fn) (bool, error) {
	return All(ctx, is, predicate.New(p))
}

func Apply(i interface{}, bf bifunction.B) function.F {
	return function.New(func(ctx context.Context, j interface{}) (interface{}, error) {
		return bf.Call(ctx, i, j)
	})
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

func Add(a interface{}) function.F {
	return Apply(a, Sum())
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
				return 0, errors.New("types not comparable")
			}
		})
}

func Sub(a interface{}) function.F {
	return Apply(a, NegativeSum())
}

func NegativeSum() bifunction.B {
	return bifunction.New(
		func(ctx context.Context, a interface{}, b interface{}) (interface{}, error) {
			if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
				return 0, errors.New("types not compatible")
			}

			switch a.(type) {
			case int:
				return a.(int) - b.(int), nil
			case int32:
				return a.(int32) - b.(int32), nil
			case int64:
				return a.(int64) - b.(int64), nil
			case uint:
				return a.(uint) - b.(uint), nil
			case uint32:
				return a.(uint32) - b.(uint32), nil
			case uint64:
				return a.(uint64) - b.(uint64), nil
			case float32:
				return a.(float32) - b.(float32), nil
			case float64:
				return a.(float64) - b.(float64), nil
			default:
				return false, errors.New("types not comparable")
			}
		})
}

func String() function.F {
	return function.New(
		func(ctx context.Context, a interface{}) (interface{}, error) {
			return fmt.Sprintf("%v", a), nil
		})
}

func Concat(sep string) bifunction.B {
	return bifunction.New(
		func(ctx context.Context, a interface{}, b interface{}) (interface{}, error) {
			as, aok := a.(string)
			if !aok {
				return "", errors.New("cannot concat non-strings")
			}

			bs, bok := b.(string)
			if !bok {
				return "", errors.New("cannot concat non-strings")
			}

			return strings.Join([]string{as, bs}, sep), nil
		})
}

func Mul(a interface{}) function.F {
	return Apply(a, Multiply())
}

func Multiply() bifunction.B {
	return bifunction.New(
		func(ctx context.Context, a interface{}, b interface{}) (interface{}, error) {
			if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
				return 0, errors.New("types not compatible")
			}

			switch a.(type) {
			case int:
				return b.(int) * a.(int), nil
			case int32:
				return b.(int32) * a.(int32), nil
			case int64:
				return b.(int64) * a.(int64), nil
			case uint:
				return b.(uint) * a.(uint), nil
			case uint32:
				return b.(uint32) * a.(uint32), nil
			case uint64:
				return b.(uint64) * a.(uint64), nil
			case float32:
				return b.(float32) * a.(float32), nil
			case float64:
				return b.(float64) * a.(float64), nil
			default:
				return false, errors.New("types not comparable")
			}
		})
}

func And(ps ...predicate.P) predicate.P {
	return predicate.New(func(ctx context.Context, a interface{}) (bool, error) {
		for _, p := range ps {
			b, err := p.Test(ctx, a)
			if err != nil {
				return false, nil
			}

			if !b {
				return false, nil
			}
		}

		return true, nil
	})
}

func Or(ps ...predicate.P) predicate.P {
	return predicate.New(func(ctx context.Context, a interface{}) (bool, error) {
		for _, p := range ps {
			b, err := p.Test(ctx, a)
			if err != nil {
				return false, nil
			}

			if b {
				return true, nil
			}
		}

		return false, nil
	})
}

func Not(p predicate.P) predicate.P {
	return predicate.New(func(ctx context.Context, a interface{}) (bool, error) {
		b, err := p.Test(ctx, a)
		return !b, err
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
	return Or(Gt(a), Eq(a))
}

func Lte(a interface{}) predicate.P {
	return Or(Lt(a), Eq(a))
}

func Neq(a interface{}) predicate.P {
	return Not(Eq(a))
}
