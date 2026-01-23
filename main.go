package main

import (
	"fmt"
	calculator "go-practice/basics/calculator"
)

func main() {
	calc := calculator.NewCalculator()

	sum := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", sum)
	fmt.Println("Calculator History:", calc.GetHistory())

	diff := calc.Subtract(10, 5)
	fmt.Printf("10 - 5 = %.2f\n", diff)
	fmt.Println("Calculator History:", calc.GetHistory())

	product := calc.Multiply(10, 5)
	fmt.Printf("10 * 5 = %.2f\n", product)
	fmt.Println("Calculator History:", calc.GetHistory())

	quotient, err := calc.Divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", quotient)
	}
	fmt.Println("Calculator History:", calc.GetHistory()) // automatic dereferencing of pointer: (*calc).GetHistory()

	quotient, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", quotient)
	}

	calc.ClearHistory()
	fmt.Println("Calculator History:", calc.GetHistory())
}
