package function

import "context"

type F interface {
	Call(ctx context.Context, i interface{}) (interface{}, error)
}

type Fn func(ctx context.Context, i interface{}) (interface{}, error)

type functionImpl struct {
	f Fn
}

func (f *functionImpl) Call(ctx context.Context, i interface{}) (interface{}, error) {
	return f.f(ctx, i)
}

func New(f Fn) F {
	return &functionImpl{f: f}
}

type multiFn struct {
	fs []F
}

func (mf *multiFn) Call(ctx context.Context, i interface{}) (interface{}, error) {
	for _, f := range mf.fs {
		var err error
		i, err = f.Call(ctx, i)
		if err != nil {
			return nil, err
		}
	}
	return i, nil
}

func Compose(fs ...F) F {
	return &multiFn{fs: fs}
}
