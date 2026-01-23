package calculator

import (
	"errors"
)

type Calculator struct {
}

func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
