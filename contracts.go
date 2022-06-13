package lep

type Expression interface {
	Equals(Expression) bool
	String() string
	Evaluate(ctx interface{}) (bool, error)
}

type Value interface {
	Expression
	Value() interface{}
}
