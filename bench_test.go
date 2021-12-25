package lab2

import (
	"fmt"
	"testing"
)

var baseInputData string = "ab+c*ef+g/+"
var cntRes string
var err error

func BenchmarkPostfixToInfix(b *testing.B) {
	const baseLength = 3000

	for i := 0; i < 20; i++ {
		input := baseInputData
		num := baseLength * (i + 1)

		for j := 0; j < num; j++ {

			input = input + baseInputData
		}

		b.Run(fmt.Sprintf("len=%d", num), func(b *testing.B) {
			cntRes, err = postfixToInfix(input)
		})
	}
}
