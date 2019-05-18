package reducer

import (
	"context"

	"github.com/samwho/fu/bifunction"
)

type R interface {
	Reduce(ctx context.Context, s interface{}, is []interface{}) (interface{}, error)
}

type Fn func(ctx context.Context, s interface{}, is []interface{}) (interface{}, error)

type bifunctionReducer struct {
	bf bifunction.B
}

func (b *bifunctionReducer) Reduce(ctx context.Context, s interface{}, is []interface{}) (interface{}, error) {
	if len(is) == 0 {
		return nil, nil
	}
	var err error
	for i := 0; i < len(is); i++ {
		s, err = b.bf.Call(ctx, s, is[i])
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func New(bf bifunction.B) R {
	return &bifunctionReducer{bf: bf}
}

func NewFn(bf bifunction.Fn) R {
	return &bifunctionReducer{bf: bifunction.New(bf)}
}
