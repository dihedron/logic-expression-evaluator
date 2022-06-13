package lep

import "testing"

func TestAnd(t *testing.T) {
	expression := `a=false && b=false`

	result, err := ParseExpression(expression)
	if err != nil {
		t.Fatal(err)
	}

	type data struct {
		a bool
		b bool
	}

	result.Evaluate(data{
		a: true,
		b: true,
	})
}
