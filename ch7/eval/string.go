package eval

import "strconv"

func (v Var) String() string {
	return string(v)
}

func (n literal) String() string {
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return "(" + b.x.String() + ")" + string(b.op) + "(" + b.y.String() + ")"
}

func (c call) String() string {
	if len(c.args) == 0 {
		return c.fn + "()"
	}
	args := c.args[0].String()
	for _, exp := range c.args[1:] {
		args += ", " + exp.String()
	}
	return c.fn + "(" + args + ")"
}
