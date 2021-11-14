package lab2

import (
	"bytes"
	"strings"
	. "gopkg.in/check.v1"
)

func Test1(c *C) {
	testInput := "+ 5 * - 4 2 3"
	var expectedOutput = "4"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	t.Assert(err, IsNil)
	t.Assert(output.String(), Equals, expectedOutput)
}

func Test2(t *testing.T) {
	testInput := "ab+c*ef+g/+"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	t.Assert(err, NotNil)
	t.Assert(output.String(), Equals, "")
}

func Test3(t *testing.T) {
	testInput := "x y ^ 5 z * / 10 +"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	t.Assert(err, NotNil)
	t.Assert(output.String(), Equals, "")
}

func Test4(t *testing.T) {
	testInput := "+ ^ - a b"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	t.Assert(err, NotNil)
	t.Assert(output.String(), Equals, "")
}
