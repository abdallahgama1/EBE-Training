package main

import "fmt"

type Record struct {
	Expr   string
	Result float64
	Err    string // empty if none
}

var history []Record

func AddHistory(expr string, result float64, err error) {
	rec := Record{Expr: expr, Result: result}
	if err != nil {
		rec.Err = err.Error()
	}
	history = append(history, rec)
}

func PrintHistory() {
	if len(history) == 0 {
		fmt.Println("No history yet.")
		return
	}
	for i, r := range history {
		if r.Err != "" {
			fmt.Printf("%d) %s => ERROR: %s\n", i, r.Expr, r.Err)
		} else {
			fmt.Printf("%d) %s = %g\n", i, r.Expr, r.Result)
		}
	}
}

func ClearHistory() {
	history = history[:0]
	fmt.Println("History cleared.")
}

func GetHistoryExpr(i int) (string, bool) {
	if i < 0 || i >= len(history) {
		return "", false
	}
	return history[i].Expr, true
}
