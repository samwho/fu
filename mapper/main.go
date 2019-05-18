package mapper

import (
	"context"

	"github.com/samwho/funcutil/function"
)

type M interface {
	Map(ctx context.Context, i []interface{}) ([]interface{}, error)
}

type Fn func(ctx context.Context, is []interface{}) ([]interface{}, error)

type functionMapper struct {
	f function.F
}

func (f *functionMapper) Map(ctx context.Context, is []interface{}) ([]interface{}, error) {
	ret := make([]interface{}, len(is))
	for idx, i := range is {
		var err error
		ret[idx], err = f.f.Call(ctx, i)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func New(f function.F) M {
	return &functionMapper{f: f}
}

func NewFn(f function.Fn) M {
	return &functionMapper{f: function.New(f)}
}
