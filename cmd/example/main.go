package main

import (
	"flag"
	"fmt"
	"go_lab_second"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputfile = flag.String("f", "", "Path/to/expression-file")
	outputFile = flag.String("o", "", "Path/to/output-file")
)

func main() {
	flag.Parse()

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
