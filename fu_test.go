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

func TestMap(t *testing.T) {
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
		{
			desc: "mul",
			f:    Mul(2),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{0, 2, 4, 6},
		},
		{
			desc: "string",
			f:    String(),
			in:   []interface{}{0, 1.1, "hello"},
			out:  []interface{}{"0", "1.1", "hello"},
		},
		{
			desc: "add 1, string",
			f:    function.Compose(Add(1), String()),
			in:   []interface{}{0, 1, 2, 3},
			out:  []interface{}{"1", "2", "3", "4"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mapped, err := Map(ctx, tC.f, tC.in)
			require.NoError(t, err)
			assert.ElementsMatch(t, tC.out, mapped)
		})
	}
}

func TestReduce(t *testing.T) {
	testCases := []struct {
		desc  string
		bf    bifunction.B
		in    []interface{}
		start interface{}
		out   interface{}
	}{
		{
			desc:  "sum",
			bf:    Sum(),
			in:    []interface{}{0, 1, 2, 3},
			start: 0,
			out:   6,
		},
		{
			desc:  "negative sum",
			bf:    NegativeSum(),
			in:    []interface{}{0, 1, 2, 3},
			start: 0,
			out:   -6,
		},
		{
			desc:  "multiply",
			bf:    Multiply(),
			in:    []interface{}{1, 2, 3, 4},
			start: 1,
			out:   24,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reduced, err := Reduce(ctx, tC.bf, tC.start, tC.in)
			require.NoError(t, err)
			assert.Equal(t, tC.out, reduced)
		})
	}
}

func TestSelect(t *testing.T) {
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
		t.Run(tC.desc, func(t *testing.T) {
			selected, err := Select(ctx, tC.p, tC.in)
			require.NoError(t, err)
			assert.ElementsMatch(t, selected, tC.out)
		})
	}
}

func TestReject(t *testing.T) {
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
		t.Run(tC.desc, func(t *testing.T) {
			selected, err := Reject(ctx, tC.p, tC.in)
			require.NoError(t, err)
			assert.ElementsMatch(t, selected, tC.out)
		})
	}
}

func TestGroupBy(t *testing.T) {
	type record struct {
		id   int
		data string
	}

	rs := []interface{}{
		record{id: 1, data: "hello"},
		record{id: 2, data: "world"},
	}

	f := function.New(func(ctx context.Context, i interface{}) (interface{}, error) {
		return i.(record).id, nil
	})

	m, err := GroupBy(ctx, f, rs)
	require.NoError(t, err)

	assert.Equal(t, m[1], rs[0])
	assert.Equal(t, m[2], rs[1])
}
