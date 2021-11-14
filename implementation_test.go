package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MySuite struct{}

func TestPostfixToInfix(t *testing.T) {
	res, err := postfixToInfix("42-3*5+")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((4-2)*3)+5)", res)
	}
}

func TestPostfixToInfix1(t *testing.T) {
	res, err := postfixToInfix("ab+c*ef+g/+")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((a+b)*c)+((e+f)/g))", res)
	}
}

func TestPostfixToInfix2(t *testing.T) {
	res, err := postfixToInfix("xy^5z*/10+")
	if assert.Nil(t, err) {
		assert.Equal(t, "(1+0)", res)
	}
}
