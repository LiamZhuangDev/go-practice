package calculator

import (
	"errors"
	"fmt"
)

type Calculator struct {
	history []string
}

func NewCalculator() *Calculator {
	return &Calculator{history: []string{}}
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

func (c *Calculator) Sum(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	c.history = append(c.history, fmt.Sprintf("Sum of %v = %.2f", numbers, total))
	return total
}

func (c *Calculator) Average(numbers ...float64) float64 {
	total := c.Sum(numbers...)
	average := total / float64(len(numbers))
	c.history = append(c.history, fmt.Sprintf("Average of %v = %.2f", numbers, average))
	return average
}

func (c Calculator) GetHistory() []string {
	return c.history
}

func (c *Calculator) ClearHistory() {
	c.history = []string{}
}
