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

type allPredicate struct {
	ps []P
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
	ps []P
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

func All(ps ...P) P {
	return &allPredicate{ps: ps}
}

func Any(ps ...P) P {
	return &anyPredicate{ps: ps}
}

type notPredicate struct {
	p P
}

func (np *notPredicate) Test(ctx context.Context, i interface{}) (bool, error) {
	b, err := np.p.Test(ctx, i)
	return !b, err
}

func Not(p P) P {
	return &notPredicate{p: p}
}
