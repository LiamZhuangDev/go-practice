package main

import (
	"fmt"
	calculator "go-practice/basics/calculator"
)

func main() {
	calc := calculator.Calculator{}

	sum := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", sum)

	diff := calc.Subtract(10, 5)
	fmt.Printf("10 - 5 = %.2f\n", diff)

	product := calc.Multiply(10, 5)
	fmt.Printf("10 * 5 = %.2f\n", product)

	quotient, err := calc.Divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", quotient)
	}

	quotient, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", quotient)
	}
}
