package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%v", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%v %c %v)", b.x, b.op, b.y)
}

func (c call) String() string {
	s := fmt.Sprintf("%s(", c.fn)
	if len(c.args) > 0 {
		s += c.args[0].String()
		for i := 1; i < len(c.args); i++ {
			s += ", " + c.args[i].String()
		}
	}
	s += ")"
	return s
}
