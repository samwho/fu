package funcutil

import "context"

type Function interface {
	Call(ctx context.Context, i interface{}) (interface{}, error)
}

type FunctionFn func(ctx context.Context, i interface{}) (interface{}, error)

type functionImpl struct {
	f FunctionFn
}

func (f *functionImpl) Call(ctx context.Context, i interface{}) (interface{}, error) {
	return f.f(ctx, i)
}

func NewFunction(f FunctionFn) Function {
	return &functionImpl{f: f}
}

type multiFn struct {
	fs []Function
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

func ComposeFunctions(fs ...Function) Function {
	return &multiFn{fs: fs}
}
