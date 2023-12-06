package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (Arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (Arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (Arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (Arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
