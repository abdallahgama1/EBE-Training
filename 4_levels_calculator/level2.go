package main

import "fmt"

// Requirements covered:
// - Store history of calculations in the session
// - Menu: (1) Calculate (2) View History (3) Clear History (4) Exit
// - Use slice with struct (num1, num2, op, result) formatted

func RunLevel2() {
	for {
		fmt.Println("\n-- Level 2: Calculator + History --")
		fmt.Println("1) Calculate")
		fmt.Println("2) View History")
		fmt.Println("3) Clear History")
		fmt.Println("4) Back")
		choice := ReadMenuInt("Choose option: ", 1, 4)

		switch choice {
		case 1:
			a := ReadFloat("Enter first number: ")
			op := ReadString("Enter operator (+, -, *, /): ")
			b := ReadFloat("Enter second number: ")

			expr := fmt.Sprintf("%g %s %g", a, op, b)
			fn, ok := OpFunc[op]
			if !ok {
				fmt.Println("Invalid operator.")
				AddHistory(expr, 0, fmt.Errorf("invalid operator %q", op))
				continue
			}
			res, err := fn(a, b)
			if err != nil {
				fmt.Println("Error:", err)
				AddHistory(expr, 0, err)
				continue
			}
			fmt.Printf("Result: %g\n", res)
			AddHistory(expr, res, nil)

		case 2:
			PrintHistory()

		case 3:
			ClearHistory()

		case 4:
			return
		}
	}
}
