package calculator

import (
	"errors"
	"fmt"
)

type Calculator struct {
	history []string
}

func (c *Calculator) Add(a, b float64) float64 {
	c.history = append(c.history, fmt.Sprintf("%.2f + %.2f = %.2f", a, b, a+b))
	return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
	c.history = append(c.history, fmt.Sprintf("%.2f - %.2f = %.2f", a, b, a-b))
	return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
	c.history = append(c.history, fmt.Sprintf("%.2f * %.2f = %.2f", a, b, a*b))
	return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	c.history = append(c.history, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, a/b))
	return a / b, nil
}

func (c Calculator) GetHistory() []string {
	return c.history
}

func (c *Calculator) ClearHistory() {
	c.history = []string{}
}
