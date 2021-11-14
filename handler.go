package lab2

import (
	"bytes"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(ch.Input)
	parsedBuffer := buffer.String()

	result, err := postfixToInfix(parsedBuffer)
	if err == nil {
		resultToByte := []byte(result)
		ch.Output.Write(resultToByte)
	} else {
		return err
	}
	return nil
}
