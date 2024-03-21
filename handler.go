package lab2

import (
	"fmt"
	"io"
	"strings"
)

type PrefixCalculatorSpy struct {
	Result string
	Error  error
}

func (prefixCalcSpy *PrefixCalculatorSpy) ConvertPrefixToInfix(expression string) (string, error) {
	return prefixCalcSpy.Result, prefixCalcSpy.Error
}

type PrefixToInfixCalculator interface {
	ConvertPrefixToInfix(expression string) (string, error)
}

type ComputeHandler struct {
	Input      io.Reader
	Output     io.Writer
	Calculator PrefixToInfixCalculator
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)

	if err != nil {
		return err
	}

	expression := strings.TrimSpace(string(data))
	result, err := ch.Calculator.ConvertPrefixToInfix(expression)

	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(fmt.Sprintf(result)))
	return err
}
