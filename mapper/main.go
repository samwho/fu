package mapper

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/samwho/fu/function"
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
	return &functionMapper{f}
}

func NewFn(f function.Fn) M {
	return &functionMapper{function.New(f)}
}

type parallelMapper struct {
	p int
	f function.F
}

type result struct {
	i int
	r interface{}
}

func (f *parallelMapper) Map(ctx context.Context, is []interface{}) ([]interface{}, error) {
	g, ctx := errgroup.WithContext(ctx)
	ret := make([]interface{}, len(is))

	idxs := make(chan int)
	c := make(chan result)

	g.Go(func() error {
		defer close(idxs)
		for idx := range is {
			select {
			case idxs <- idx:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		return nil
	})

	for j := 0; j < f.p; j++ {
		g.Go(func() error {
			for i := range idxs {
				r, err := f.f.Call(ctx, is[i])
				if err != nil {
					return err
				}
				select {
				case c <- result{i, r}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(c)
	}()

	for r := range c {
		ret[r.i] = r.r
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return ret, nil
}

func Parallel(parallelism int, f function.F) M {
	return &parallelMapper{parallelism, f}
}
