package predicate

import (
	"context"
)

type P interface {
	Test(ctx context.Context, i interface{}) (bool, error)
}

type Fn func(ctx context.Context, i interface{}) (bool, error)

type predicateImpl struct {
	f Fn
}

func (p *predicateImpl) Test(ctx context.Context, i interface{}) (bool, error) {
	return p.f(ctx, i)
}

func New(f Fn) P {
	return &predicateImpl{f: f}
}
