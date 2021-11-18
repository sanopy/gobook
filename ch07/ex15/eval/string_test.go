package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		{"1 < 2", Env{"x": 1}, "1"},
		{"1 < x", Env{"x": 1}, "0"},
		{"2 > 1", Env{"x": 1}, "1"},
		{"1 > x", Env{"x": 1}, "0"},
		{"1 < 2 ? 1 : 2", Env{"x": 2}, "1"},
		{"1 < 2 ? 1 + x : 2 * x", Env{"x": 2}, "3"},
		{"1 > 2 ? 1 : 2", Env{"x": 2}, "2"},
		{"1 > 2 ? 1 + x : 2 * x", Env{"x": 2}, "4"},
	}
	for _, test := range tests {
		fmt.Printf("\noriginal expr: %s\n", test.expr)
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		fmt.Printf(".String(): %v\n", expr)
		expr2, err := Parse(expr.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		got := fmt.Sprintf("%.6g", expr2.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
