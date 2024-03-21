package main

import (
	"fmt"
)

func isOperator(x rune) bool {
	switch x {
	case '+', '-', '*', '/', '^', '%':
		return true
	}
	return false
}

func convert(str string) string {
	stack := []string{}

	l := len(str)

	for i := l - 1; i >= 0; i-- {
		c := rune(str[i])

		if isOperator(c) {
			if len(stack) >= 2 {
				op1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				temp := fmt.Sprintf("(%s%c%s)", op1, c, op2)
				stack = append(stack, temp)
			} else {
				return "Invalid prefix expression"
			}
		} else {
			stack = append(stack, string(c))
		}
	}
	return stack[len(stack)-1]
}

func main() {
	exp := "your math problem here"

	fmt.Println("Infix :", convert(exp))
}
