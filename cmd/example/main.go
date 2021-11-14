package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	lab2 "github.com/software-engineering-components/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	fileExpression = flag.String("f", "", "Expression to read from file")
	outputExpression = flag.String("o", "", "Output Expression")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer

	if (len(*inputExpression) == 0 && len(*fileExpression) == 0) ||
			(len(*inputExpression) > 0 && len(*fileExpression) > 0) {
		os.Stderr.WriteString("Invalid Input")
		return
	}

	if len(*inputExpression) > 0 {
		input = strings.NewReader(*inputExpression)
	}

	if len(*fileExpression) > 0 {
		inputFileBytes, err := os.ReadFile(*fileExpression)
		if err != nil {
			os.Stderr.WriteString("File Error")
		}
		input = bytes.NewReader(inputFileBytes)
	}

	if len(*outputExpression) > 0 {
		outputFile, err := os.Create(*outputExpression)
		if err != nil {
			os.Stderr.WriteString("Output Error")
		}
		output = outputFile
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input: input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString("Compute handler error")
	}
}