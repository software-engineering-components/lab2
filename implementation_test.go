package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := postfixToInfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func TestPostfixToInfix(t *testing.T) {
	res, err := postfixToInfix("ab+c*ef+g/+")
	if assert.Nil(t, err) {
		assert.Equal(t, "abc*de-/+", res)
	}
}

func TestPostfixToInfix(t *testing.T) {
	res, err := postfixToInfix("x y ^ 5 z * / 10 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "x ^ y / (5 * z) + 10", res)
	}
}
