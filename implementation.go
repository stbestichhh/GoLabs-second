package lab2

import (
	"errors"
	"strings"
)

type PrefixCalculator struct{}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func (ptic *PrefixCalculator) ConvertPrefixToInfix(str string) (string, error) {
	tokens := strings.Fields(str)
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("invalid prefix expression")
			}

			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack = append(stack, "("+op1+token+op2+")")
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid prefix expression")
	}

	return stack[0], nil
}
