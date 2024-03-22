package lab2

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestConvertPrefixToInfix(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    string
		expectedErr bool
	}{
		{
			name:        "Simple expression",
			expression:  "+ 2 3",
			expected:    "(2+3)",
			expectedErr: false,
		},
		{
			name:        "Complex expression",
			expression:  "+ * 2 3 4",
			expected:    "((2*3)+4)",
			expectedErr: false,
		},
		{
			name:        "Empty expression",
			expression:  "",
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "Invalid expression",
			expression:  "+ 0",
			expected:    "",
			expectedErr: true,
		},
	}

	prefixCalc := PrefixCalculator{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := prefixCalc.ConvertPrefixToInfix(test.expression)
if test.expectedErr {
				assert.Error(t, err, "expected error")
			} else {
				assert.NoError(t, err, "unexpected error")
				assert.Equal(t, test.expected, result, "result not as expected")
			}
		})
	}
}

// ExampleConvertPrefixToInfix demonstrates how to use the ConvertPrefixToInfix
func ExampleConvertPrefixToInfix() {
	prefixCalc := PrefixCalculator{}

	// Простий приклад виразу
	result, err := prefixCalc.ConvertPrefixToInfix("+ 2 3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	// Вивід: (2+3)

	// Складний приклад виразу
	result, err = prefixCalc.ConvertPrefixToInfix("+ * 2 3 4")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	// Вивід: ((2*3)+4)

	// Порожній вираз (очікується помилка)
	result, err = prefixCalc.ConvertPrefixToInfix("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	// Вивід: Error: empty expression

	// Неправильний вираз (очікується помилка)
	result, err = prefixCalc.ConvertPrefixToInfix("+ 0")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
	// Вивід: Error: invalid expression
}