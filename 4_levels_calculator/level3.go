package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Requirements covered:
// - Input full expression (e.g., "2 + 3 * 4")
// - Precedence (*, / before +, -) using shunting-yard
// - Use maps for operator funcs/precedence
// - Invalid expression errors
// - Maintain history (reuse Level 2)
// - Replay: re-run expression from history by index

func RunLevel3() {
	for {
		fmt.Println("\n-- Level 3: Expressions with precedence --")
		fmt.Println("1) Evaluate expression")
		fmt.Println("2) View History")
		fmt.Println("3) Replay from History")
		fmt.Println("4) Clear History")
		fmt.Println("5) Back")
		choice := ReadMenuInt("Choose option: ", 1, 5)

		switch choice {
		case 1:
			line := ReadString("Enter expression (e.g., 2 + 3 * 4): ")
			res, err := EvalExpression(line)
			if err != nil {
				fmt.Println("Error:", err)
				AddHistory(line, 0, err)
			} else {
				fmt.Printf("Result: %g\n", res)
				AddHistory(line, res, nil)
			}

		case 2:
			PrintHistory()

		case 3:
			if len(history) == 0 {
				fmt.Println("No history to replay.")
				continue
			}
			fmt.Printf("Select index [0..%d]: ", len(history)-1)
			idx := ReadMenuInt("", 0, len(history)-1)
			expr, ok := GetHistoryExpr(idx)
			if !ok {
				fmt.Println("Invalid index.")
				continue
			}
			fmt.Println("Replaying:", expr)
			res, err := EvalExpression(expr)
			if err != nil {
				fmt.Println("Error:", err)
				AddHistory(expr, 0, err)
			} else {
				fmt.Printf("Result: %g\n", res)
				AddHistory(expr, res, nil)
			}

		case 4:
			ClearHistory()

		case 5:
			return
		}
	}
}

// ---- Expression parsing & evaluation (no parentheses) ----

// Tokenize numbers (floats) and operators + - * /
func tokenize(s string) ([]string, error) {
	var tokens []string
	i := 0
	for i < len(s) {
		r := rune(s[i])
		if unicode.IsSpace(r) {
			i++
			continue
		}
		// number (digits and at most one '.')
		if unicode.IsDigit(r) || r == '.' {
			start := i
			dot := r == '.'
			i++
			for i < len(s) {
				r2 := rune(s[i])
				if unicode.IsDigit(r2) {
					i++
					continue
				}
				if r2 == '.' && !dot {
					dot = true
					i++
					continue
				}
				break
			}
			tokens = append(tokens, s[start:i])
			continue
		}
		// operator
		switch r {
		case '+', '-', '*', '/':
			tokens = append(tokens, string(r))
			i++
		default:
			return nil, fmt.Errorf("invalid character %q", r)
		}
	}
	return tokens, nil
}

// Convert infix → RPN (Reverse Polish Notation) via shunting-yard (no parens)
func infixToRPN(tokens []string) ([]string, error) {
	var out []string
	var ops []string

	// simple validation: token sequence must alternate number/op roughly
	expectNumber := true
	for _, t := range tokens {
		if isNumber(t) {
			if !expectNumber {
				return nil, errors.New("unexpected number")
			}
			out = append(out, t)
			expectNumber = false
			continue
		}
		// operator
		if _, ok := OpPrec[t]; !ok {
			return nil, fmt.Errorf("unknown operator %q", t)
		}
		if expectNumber {
			return nil, errors.New("expression cannot start with an operator or have two operators in a row")
		}
		// precedence handling
		for len(ops) > 0 {
			top := ops[len(ops)-1]
			if OpPrec[top] >= OpPrec[t] {
				out = append(out, top)
				ops = ops[:len(ops)-1]
			} else {
				break
			}
		}
		ops = append(ops, t)
		expectNumber = true
	}
	if expectNumber {
		return nil, errors.New("expression cannot end with an operator")
	}
	for i := len(ops) - 1; i >= 0; i-- {
		out = append(out, ops[i])
	}
	return out, nil
}

func evalRPN(rpn []string) (float64, error) {
	var st []float64
	for _, t := range rpn {
		if isNumber(t) {
			v, _ := strconv.ParseFloat(t, 64)
			st = append(st, v)
			continue
		}
		// operator: pop b, a
		if len(st) < 2 {
			return 0, errors.New("malformed expression")
		}
		b := st[len(st)-1]
		a := st[len(st)-2]
		st = st[:len(st)-2]

		fn := OpFunc[t]
		res, err := fn(a, b)
		if err != nil {
			return 0, err
		}
		st = append(st, res)
	}
	if len(st) != 1 {
		return 0, errors.New("malformed expression")
	}
	return st[0], nil
}

func EvalExpression(s string) (float64, error) {
	toks, err := tokenize(s)
	if err != nil {
		return 0, err
	}
	rpn, err := infixToRPN(toks)
	if err != nil {
		return 0, err
	}
	return evalRPN(rpn)
}

func isNumber(s string) bool {
	if s == "" {
		return false
	}
	// quick check: first rune digit or dot
	r := rune(s[0])
	return unicode.IsDigit(r) || r == '.'
}
