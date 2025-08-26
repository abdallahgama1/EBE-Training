package main

import "fmt"

func main() {
	for {
		fmt.Println("\n === Calculator === ")
		fmt.Println("1) level 1 Basic Calculator")
		fmt.Println("2) level 2 Basic Calculator + History ")
		fmt.Println("3) level 3 Expressions with Precedence + History + Replay ")

		choice := ReadMenuInt("Choose option: ", 1, 4)

		switch choice {
		case 1:
			RunLevel1()
		case 2:
			RunLevel2()
		case 3:
			RunLevel3()
		case 4:
			fmt.Println("Goodbye")
			return
		}
	}
}
