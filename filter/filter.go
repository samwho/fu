package filter

import (
	"context"

	"github.com/samwho/fu/predicate"
)

type F interface {
	Filter(ctx context.Context, is []interface{}) ([]interface{}, error)
}

type predicateFilter struct {
	p predicate.P
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
	fs []F
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

func NewFn(f predicate.Fn) F {
	return &predicateFilter{p: predicate.New(f)}
}

func New(p predicate.P) F {
	return &predicateFilter{p: p}
}
