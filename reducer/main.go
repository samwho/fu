package reducer

import (
	"context"

	"github.com/samwho/fu/bifunction"
)

type R interface {
	Reduce(ctx context.Context, is []interface{}) (interface{}, error)
}

type Fn func(ctx context.Context, is []interface{}) (interface{}, error)

type bifunctionReducer struct {
	bf bifunction.B
}

func (b *bifunctionReducer) Reduce(ctx context.Context, is []interface{}) (interface{}, error) {
	if len(is) == 0 {
		return nil, nil
	}
	r := is[0]
	var err error
	for i := 1; i < len(is); i++ {
		r, err = b.bf.Call(ctx, r, is[i])
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func New(bf bifunction.B) R {
	return &bifunctionReducer{bf: bf}
}

func NewFn(bf bifunction.Fn) R {
	return &bifunctionReducer{bf: bifunction.New(bf)}
}
