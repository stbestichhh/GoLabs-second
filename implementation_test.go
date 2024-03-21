package lab2

import (
	"testing"
	"errors"
	//lab2 "go_lab_second"
	"github.com/stretchr/testify/assert"
)

func TestConvertPrefixToInfix(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   string
		err        error
	}{
		{
			name:       "Simple expression",
			expression: "+ 2 3",
			expected:   "2 + 3",
			err:        nil,
		},
		{
			name:       "Complex expression",
			expression: "+ * 2 3 4",
			expected:   "2 * 3 + 4",
			err:        nil,
		},
		{
			name:       "Empty expression",
			expression: "",
			expected:   "",
			err:        errors.New("empty expression"),
		},
		{
			name:       "Invalid expression",
			expression: "abc",
			expected:   "",
			err:        errors.New("invalid expression"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := PrefixCalculatorSpy{Result: tt.expected, Error: tt.err}
			result, err := calc.ConvertPrefixToInfix(tt.expression)

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func ExampleConvertPrefixToInfix() {}
