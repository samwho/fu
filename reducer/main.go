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
	if len(is) == 1 {
		return is[0], nil // TODO(samwho): unsure about this, revisit
	}

	var err error
	ret := is[0]
	for i := 1; i < len(is); i++ {
		ret, err = b.bf.Call(ctx, ret, is[i])
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func New(bf bifunction.B) R {
	return &bifunctionReducer{bf: bf}
}

func NewFn(bf bifunction.Fn) R {
	return &bifunctionReducer{bf: bifunction.New(bf)}
}
