package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []string{
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"5 / 9 * (F - 32)"}

	for _, test := range tests {
		expr1, err := Parse(test)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		fmt.Printf("test = %s  string = %s \n", test, expr1.String())

	}
}
