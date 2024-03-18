package main

import (
	"flag"
	"fmt"
	"go_lab_second"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Path/to/expression-file")
	outputFile = flag.String("o", "", "Path/to/output-file")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)

		if err != nil {
			fmt.Println(os.Stderr, "Cannot open file.", err)
		}

		defer file.Close()

		input = file;
	} else {
		fmt.Println(os.Stderr, "Usage: go run cmd/example/main.go -e <expression> | -f <path/to/expression-file [-o <path/to/outputfile>]")
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixToInfix("+ 2 2")
	fmt.Println(res)
}
