package main

import "fmt"

// Requirements covered:
// - Read two float64 numbers and an operator
// - Perform + - * /
// - Handle division by zero
// - Option to exit (back to main menu)

func RunLevel1() {
	fmt.Println("\n-- Level 1: Basic Calculator --")

	a := ReadFloat("Enter first number: ")
	op := ReadString("Enter operator (+, -, *, /): ")
	b := ReadFloat("Enter second number: ")

	var (
		res float64
		err error
	)

	switch op {
	case "+", "-", "*", "/":
		res, err = OpFunc[op](a, b)
	default:
		fmt.Println("Invalid operator.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %g\n", res)
}
