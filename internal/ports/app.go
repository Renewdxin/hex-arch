package ports

// api allows the outside world to interact with the application that implements the interface
type APIPort interface {
	GetAddition(a int32, b int32) (int32, error)
	GetSubtraction(a int32, b int32) (int32, error)
	GetMultiplication(a int32, b int32) (int32, error)
	GetDivision(a int32, b int32) (int32, error)
}
