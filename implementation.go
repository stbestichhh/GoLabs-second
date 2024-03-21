package lab2

import (
	"fmt"
)

type PrefixCalculator struct {}

func isOperator(x rune) bool {
	switch x {
	case '+', '-', '*', '/', '^', '%':
		return true
	}
	return false
}

func (ptic *PrefixCalculator) ConvertPrefixToInfix(str string) (string, error) {
	stack := []string{}

	length:= len(str)

	for i := length - 1; i >= 0; i-- {
		c := rune(str[i])

		if isOperator(c) {
			if len(stack) >= 2 {
				op1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				op2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				temp := fmt.Sprintf("%s%c%s", op1, c, op2)
				stack = append(stack, temp)
			} else {
				return "", fmt.Errorf("Invalid prefix expression");
			}
		} else {
			stack = append(stack, string(c))
		}
	}

	if (len(stack) != 1) {
		return "", fmt.Errorf("Invalid prefix expression")
	}

	return stack[len(stack)-1], nil;
}
