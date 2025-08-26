package main

import "errors"

func Add(a, b float64) (float64, error) { return a + b, nil }
func Sub(a, b float64) (float64, error) { return a - b, nil }
func Mul(a, b float64) (float64, error) { return a * b, nil }
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// For Level 3: a small registry (map) to help precedence/eval
var OpFunc = map[string]func(float64, float64) (float64, error){
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

var OpPrec = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
