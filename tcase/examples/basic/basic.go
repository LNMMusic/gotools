package basic

import "errors"

var (
	ErrDivideByZero = errors.New("error: divide by zero")
)

func Sum(a, b int) (result int) {
	result = a + b
	return
}

func Div(a, b float64) (result float64, err error) {
	if b == 0 {
		err = ErrDivideByZero
		return
	}

	result = a / b
	return
}