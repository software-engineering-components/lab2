package lab2

import (
	"bytes"
	"strings"
	. "gopkg.in/check.v1"
)

func (s *MySuite) Test12(c *C) {
	testInput := "42-3*5+"
	var expectedOutput = "(((4-2)*3)+5)"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	c.Assert(err, IsNil)
	c.Assert(output.String(), Equals, expectedOutput)
}

func (s *MySuite) Test13(c *C) {
	testInput := "ab+c*ef+g/+"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	c.Assert(err, NotNil)
	c.Assert(output.String(), Equals, "")
}

func (s *MySuite) Test14(c *C) {
	testInput := "xy^5z*/10+"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	c.Assert(err, NotNil)
	c.Assert(output.String(), Equals, "")
}

func (s *MySuite) Test15(c *C) {
	testInput := "+ ^ - a b"

	input := strings.NewReader(testInput)
	output := bytes.NewBufferString("")
	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	c.Assert(err, NotNil)
	c.Assert(output.String(), Equals, "")
}
