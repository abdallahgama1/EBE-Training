package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

func ReadString(prompt string) string {
	for {
		fmt.Print(prompt)
		s, err := stdin.ReadString('\n')
		if err == nil {
			s = strings.TrimSpace(s)
			if s != "" {
				return s
			}
		}
		fmt.Println("Invalid input, try again.")
	}
}

func ReadFloat(prompt string) float64 {
	for {
		raw := ReadString(prompt)
		if v, err := strconv.ParseFloat(raw, 64); err == nil {
			return v
		}
		fmt.Println("Please enter a valid number (e.g. 12, 3.14).")
	}
}

func ReadMenuInt(prompt string, min, max int) int {
	for {
		raw := ReadString(prompt)
		n, err := strconv.Atoi(raw)
		if err == nil && n >= min && n <= max {
			return n
		}
		fmt.Printf("Enter a number between %d and %d.\n", min, max)
	}
}
