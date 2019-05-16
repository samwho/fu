package funcutil

/*
import (
	"context"
)

type Filter interface {
	Filter(ctx context.Context, is []interface{}) ([]interface{}, error)
}

type predicateFilter struct {
	p Predicate
}

func (pf *predicateFilter) Filter(ctx context.Context, is []interface{}) ([]interface{}, error) {
	var filtered []interface{}
	for _, i := range is {
		b, err := pf.p.Test(ctx, i)
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

type multiFilter struct {
	fs []Filter
}

func (mf *multiFilter) Filter(ctx context.Context, is []interface{}) ([]interface{}, error) {
	for _, f := range mf.fs {
		var err error
		is, err = f.Filter(ctx, is)
		if err != nil || len(is) == 0 {
			return []interface{}{}, err
		}
	}
	return is, nil
}

func NewFilterFn(f PredicateFn) Filter {
	return &predicateFilter{p: NewPredicate(f)}
}

func NewFilter(p Predicate) Filter {
	return &predicateFilter{p: p}
}
*/
