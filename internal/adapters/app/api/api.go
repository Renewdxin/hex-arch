package api

import "hex-arch/internal/ports"

type Adapter struct {
	// depedencies injection
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}

func (api Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	err = api.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := api.arith.Subtraction(a, b)
	err = api.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	return api.arith.Multiplication(a, b)
}

func (api Adapter) GetDivision(a int32, b int32) (int32, error) {
	return api.arith.Division(a, b)
}
