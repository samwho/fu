package fu

import (
	"context"
	"testing"

	"github.com/samwho/fu/predicate"

	"github.com/samwho/fu/bifunction"

	"github.com/samwho/fu/function"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func TestBifunctions(t *testing.T) {
	testCases := []struct {
		bf          bifunction.B
		i           interface{}
		j           interface{}
		expectedRes interface{}
		expectedErr bool
	}{
		{bf: Multiply(), i: int(3), j: int(3), expectedRes: int(9)},
		{bf: Multiply(), i: int32(3), j: int32(3), expectedRes: int32(9)},
		{bf: Multiply(), i: int64(3), j: int64(3), expectedRes: int64(9)},
		{bf: Multiply(), i: uint(3), j: uint(3), expectedRes: uint(9)},
		{bf: Multiply(), i: uint32(3), j: uint32(3), expectedRes: uint32(9)},
		{bf: Multiply(), i: uint64(3), j: uint64(3), expectedRes: uint64(9)},
		{bf: Multiply(), i: float32(3), j: float32(3), expectedRes: float32(9)},
		{bf: Multiply(), i: float64(3), j: float64(3), expectedRes: float64(9)},
		{bf: Multiply(), i: int(3), j: float64(3), expectedErr: true}, // could reasonably not be an error
		{bf: Multiply(), i: "hello", j: "world", expectedErr: true},
		{bf: Multiply(), i: struct{}{}, j: struct{}{}, expectedErr: true},

		{bf: Sum(), i: int(3), j: int(3), expectedRes: int(6)},
		{bf: Sum(), i: int32(3), j: int32(3), expectedRes: int32(6)},
		{bf: Sum(), i: int64(3), j: int64(3), expectedRes: int64(6)},
		{bf: Sum(), i: uint(3), j: uint(3), expectedRes: uint(6)},
		{bf: Sum(), i: uint32(3), j: uint32(3), expectedRes: uint32(6)},
		{bf: Sum(), i: uint64(3), j: uint64(3), expectedRes: uint64(6)},
		{bf: Sum(), i: float32(3), j: float32(3), expectedRes: float32(6)},
		{bf: Sum(), i: float64(3), j: float64(3), expectedRes: float64(6)},
		{bf: Sum(), i: int(3), j: float64(3), expectedErr: true}, // could reasonably not be an error
		{bf: Sum(), i: "hello", j: "world", expectedErr: true},
		{bf: Sum(), i: struct{}{}, j: struct{}{}, expectedErr: true},

		{bf: NegativeSum(), i: int(3), j: int(2), expectedRes: int(1)},
		{bf: NegativeSum(), i: int32(3), j: int32(2), expectedRes: int32(1)},
		{bf: NegativeSum(), i: int64(3), j: int64(2), expectedRes: int64(1)},
		{bf: NegativeSum(), i: uint(3), j: uint(2), expectedRes: uint(1)},
		{bf: NegativeSum(), i: uint32(3), j: uint32(2), expectedRes: uint32(1)},
		{bf: NegativeSum(), i: uint64(3), j: uint64(2), expectedRes: uint64(1)},
		{bf: NegativeSum(), i: float32(3), j: float32(2), expectedRes: float32(1)},
		{bf: NegativeSum(), i: float64(3), j: float64(2), expectedRes: float64(1)},
		{bf: NegativeSum(), i: int(3), j: float64(2), expectedErr: true}, // could reasonably not be an error
		{bf: NegativeSum(), i: "hello", j: "world", expectedErr: true},
		{bf: NegativeSum(), i: struct{}{}, j: struct{}{}, expectedErr: true},

		{bf: Join(", "), i: "hello", j: "world", expectedRes: "hello, world"},
		{bf: Join(", "), i: "hello", j: 1, expectedErr: true},
		{bf: Join(", "), i: 1, j: 2, expectedErr: true},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run("", func(t *testing.T) {
			t.Parallel()
			res, err := tC.bf.Call(ctx, tC.i, tC.j)
			if tC.expectedErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tC.expectedRes, res)
			}
		})
	}
}

func TestFunctions(t *testing.T) {
	testCases := []struct {
		f           function.F
		in          interface{}
		out         interface{}
		expectedErr bool
	}{
		{f: Add(int(1)), in: int(3), out: int(4)},
		{f: Add(int32(1)), in: int32(3), out: int32(4)},
		{f: Add(int64(1)), in: int64(3), out: int64(4)},
		{f: Add(uint(1)), in: uint(3), out: uint(4)},
		{f: Add(uint32(1)), in: uint32(3), out: uint32(4)},
		{f: Add(uint64(1)), in: uint64(3), out: uint64(4)},
		{f: Add(float32(1)), in: float32(3), out: float32(4)},
		{f: Add(float64(1)), in: float64(3), out: float64(4)},
		{f: Add(int(1)), in: float64(3), expectedErr: true},

		{f: Mul(int(2)), in: int(3), out: int(6)},
		{f: Mul(int32(2)), in: int32(3), out: int32(6)},
		{f: Mul(int64(2)), in: int64(3), out: int64(6)},
		{f: Mul(uint(2)), in: uint(3), out: uint(6)},
		{f: Mul(uint32(2)), in: uint32(3), out: uint32(6)},
		{f: Mul(uint64(2)), in: uint64(3), out: uint64(6)},
		{f: Mul(float32(2)), in: float32(3), out: float32(6)},
		{f: Mul(float64(2)), in: float64(3), out: float64(6)},
		{f: Mul(int(2)), in: float64(3), expectedErr: true},

		{f: Sub(int(2)), in: int(3), out: int(1)},
		{f: Sub(int32(2)), in: int32(3), out: int32(1)},
		{f: Sub(int64(2)), in: int64(3), out: int64(1)},
		{f: Sub(uint(2)), in: uint(3), out: uint(1)},
		{f: Sub(uint32(2)), in: uint32(3), out: uint32(1)},
		{f: Sub(uint64(2)), in: uint64(3), out: uint64(1)},
		{f: Sub(float32(2)), in: float32(3), out: float32(1)},
		{f: Sub(float64(2)), in: float64(3), out: float64(1)},
		{f: Sub(int(2)), in: float64(3), expectedErr: true},

		{f: String(), in: int(2), out: "2"},
		{f: String(), in: int32(2), out: "2"},
		{f: String(), in: int64(2), out: "2"},
		{f: String(), in: uint(2), out: "2"},
		{f: String(), in: uint32(2), out: "2"},
		{f: String(), in: uint64(2), out: "2"},
		{f: String(), in: float32(2), out: "2"},
		{f: String(), in: float64(2), out: "2"},
		{f: String(), in: "2", out: "2"},
		{f: String(), in: struct{ i int }{2}, out: "{2}"},

		{f: Field("A"), in: struct{ A int }{2}, out: 2},
		{f: Field("B"), in: struct{ B interface{} }{}, out: nil},
		{f: Field("B"), in: struct{ A int }{}, expectedErr: true},
		{f: Field(""), in: struct{ A int }{}, expectedErr: true},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run("", func(t *testing.T) {
			t.Parallel()
			res, err := tC.f.Call(ctx, tC.in)
			if tC.expectedErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tC.out, res)
			}
		})
	}
}
func TestMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		f    function.F
		in   []interface{}
		out  []interface{}
	}{
		{
			desc: "add",
			f:    Add(1),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{1, 2, 3, 4},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			mapped, err := Map(ctx, tC.in, tC.f)
			require.NoError(t, err)
			assert.ElementsMatch(t, tC.out, mapped)
		})
	}
}

func TestParallelMap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		f    function.F
		in   []interface{}
		out  []interface{}
	}{
		{
			desc: "add",
			f:    Add(1),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{1, 2, 3, 4},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			mapped, err := ParallelMap(ctx, 16, tC.in, tC.f)
			require.NoError(t, err)
			assert.ElementsMatch(t, tC.out, mapped)
		})
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		bf   bifunction.B
		in   []interface{}
		out  interface{}
	}{
		{
			desc: "sum",
			bf:   Sum(),
			in:   []interface{}{0, 1, 2, 3},
			out:  6,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			reduced, err := Reduce(ctx, tC.in, tC.bf)
			require.NoError(t, err)
			assert.Equal(t, tC.out, reduced)
		})
	}
}

func TestSelect(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		p    predicate.P
		in   []interface{}
		out  []interface{}
	}{
		{
			desc: "greater than 2",
			p:    Gt(2),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{3},
		},
		{
			desc: "greater than 2, less than 5",
			p:    And(Gt(2), Lt(5)),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{3, 4},
		},
		{
			desc: "less than or equal to 2",
			p:    Lte(2),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{0, 1, 2},
		},
		{
			desc: "greater than or equal to 2",
			p:    Gte(2),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{2, 3, 4, 5, 6, 7},
		},
		{
			desc: "not (greater than 2, less than 5)",
			p:    Not(And(Gt(2), Lt(5))),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{0, 1, 2, 5, 6, 7},
		},
		{
			desc: "strings",
			p:    Gt("c"),
			in:   []interface{}{"a", "b", "c", "d", "e"},
			out:  []interface{}{"d", "e"},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			selected, err := Select(ctx, tC.in, tC.p)
			require.NoError(t, err)
			assert.ElementsMatch(t, selected, tC.out)
		})
	}
}

func TestReject(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc string
		p    predicate.P
		in   []interface{}
		out  []interface{}
	}{
		{
			desc: "greater than 2",
			p:    Gt(2),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{0, 1, 2},
		},
		{
			desc: "greater than 2, less than 5",
			p:    And(Gt(2), Lt(5)),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{0, 1, 2, 5, 6, 7},
		},
		{
			desc: "not (greater than 2, less than 5)",
			p:    Not(And(Gt(2), Lt(5))),
			in:   []interface{}{0, 1, 2, 3, 4, 5, 6, 7},
			out:  []interface{}{3, 4},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			selected, err := Reject(ctx, tC.in, tC.p)
			require.NoError(t, err)
			assert.ElementsMatch(t, selected, tC.out)
		})
	}
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	type record struct {
		ID   int
		Data string
	}

	rs := []interface{}{
		record{ID: 1, Data: "hello"},
		record{ID: 2, Data: "world"},
	}

	m, err := GroupBy(ctx, Field("ID"), rs)
	require.NoError(t, err)

	assert.Equal(t, rs[0], m[1][0])
	assert.Equal(t, rs[1], m[2][0])
}
