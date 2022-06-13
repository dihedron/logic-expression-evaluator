package lep

type Expression interface {
	Equals(Expression) bool
	Evaluate(ctx interface{}) (bool, error)
	String() string
}

type Value interface {
	Expression
	Value() interface{}
}
