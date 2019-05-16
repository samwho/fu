package funcutil

import (
	"context"
	"errors"
	"reflect"
)

type Predicate interface {
	Test(ctx context.Context, i interface{}) (bool, error)
}

type PredicateFn func(ctx context.Context, i interface{}) (bool, error)

type predicateImpl struct {
	f PredicateFn
}

func (p *predicateImpl) Test(ctx context.Context, i interface{}) (bool, error) {
	return p.f(ctx, i)
}

func NewPredicate(f PredicateFn) Predicate {
	return &predicateImpl{f: f}
}

type allPredicate struct {
	ps []Predicate
}

func (mp *allPredicate) Test(ctx context.Context, i interface{}) (bool, error) {
	for _, p := range mp.ps {
		b, err := p.Test(ctx, i)
		if err != nil || !b {
			return false, err
		}
	}
	return true, nil
}

type anyPredicate struct {
	ps []Predicate
}

func (mp *anyPredicate) Test(ctx context.Context, i interface{}) (bool, error) {
	for _, p := range mp.ps {
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

func All(ps ...Predicate) Predicate {
	return &allPredicate{ps: ps}
}

func Any(ps ...Predicate) Predicate {
	return &anyPredicate{ps: ps}
}

func Gt(a interface{}) Predicate {
	return NewPredicate(func(ctx context.Context, b interface{}) (bool, error) {
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

func Lt(a interface{}) Predicate {
	return NewPredicate(func(ctx context.Context, b interface{}) (bool, error) {
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

func Eq(a interface{}) Predicate {
	return NewPredicate(func(ctx context.Context, b interface{}) (bool, error) {
		return reflect.DeepEqual(a, b), nil
	})
}

type notPredicate struct {
	p Predicate
}

func (np *notPredicate) Test(ctx context.Context, i interface{}) (bool, error) {
	b, err := np.p.Test(ctx, i)
	return !b, err
}

func Not(p Predicate) Predicate {
	return &notPredicate{p: p}
}

func Gte(a interface{}) Predicate {
	return Any(Gt(a), Eq(a))
}

func Lte(a interface{}) Predicate {
	return Any(Lt(a), Eq(a))
}

func Neq(a interface{}) Predicate {
	return Not(Eq(a))
}
