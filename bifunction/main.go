package bifunction

import (
	"context"
)

type B interface {
	Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error)
}

type Fn func(ctx context.Context, i interface{}, j interface{}) (interface{}, error)

type bifunctionImpl struct {
	bf Fn
}

func (bf *bifunctionImpl) Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
	return bf.bf(ctx, i, j)
}

func New(bf Fn) B {
	return &bifunctionImpl{bf: bf}
}

type multiBiFn struct {
	bfs []B
}

func (mbf *multiBiFn) Call(ctx context.Context, i interface{}, j interface{}) (interface{}, error) {
	for _, bf := range mbf.bfs {
		var err error
		i, err = bf.Call(ctx, i, j)
		if err != nil {
			return nil, err
		}
	}
	return i, nil
}
