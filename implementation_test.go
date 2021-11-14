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

func ExamplePostfixToInfix() {
	res, _ := postfixToInfix("ab+c*ef+g/+")
	r, _ := postfixToInfix("+ 5 * - 4 2 3")

	fmt.Println(res, r)
}
